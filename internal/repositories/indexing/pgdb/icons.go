package pgdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"iconrepo/internal/app/domain"
	"iconrepo/internal/repositories/indexing"

	"github.com/jackc/pgconn"
	"github.com/rs/zerolog"
)

var (
	ErrDuplicateRows = errors.New("duplicat rows")
)

func MapDBError(err error) error {
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return nil
	}
	fmt.Printf(">>>>> MapDBError: ok %v\n", pgErr.Code)
	if pgErr.Code == "23505" {
		return ErrDuplicateRows
	}
	if pgErr.Code == "42P01" {
		return indexing.ErrTableNotFound
	}
	return nil
}

func IsDBError(err error, target error) bool {
	if knownDBError := MapDBError(err); errors.Is(knownDBError, target) {
		return true
	}
	return false
}

func describeIconInTx(tx *sql.Tx, iconName string, forUpdate bool) (domain.IconDescriptor, error) {
	var err error
	var rows *sql.Rows

	var forUpdateClause = ""
	if forUpdate {
		forUpdateClause = " FOR UPDATE"
	}
	var iconSQL = "SELECT id, modified_by FROM icon WHERE name = $1" + forUpdateClause
	var iconfilesSQL = "SELECT file_format, icon_size FROM icon_file " +
		"WHERE icon_id = $1 " +
		"ORDER BY file_format, icon_size" + forUpdateClause
	var tagsSQL = "SELECT text FROM tag, icon_to_tags " +
		"WHERE icon_to_tags.icon_id = $1 " +
		"AND icon_to_tags.tag_id = tag.id" + forUpdateClause

	var iconId int
	var modifiedBy string
	err = tx.QueryRow(iconSQL, iconName).Scan(&iconId, &modifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.IconDescriptor{}, fmt.Errorf("icon %s not found: %w", iconName, domain.ErrIconNotFound)
		} else {
			return domain.IconDescriptor{}, fmt.Errorf("error while retrieving icon '%s' from database: %w", iconName, err)
		}
	}

	iconfiles := make([]domain.IconfileDescriptor, 0, 10)
	emptyIcon := domain.IconDescriptor{}
	err = func() error {
		rows, err = tx.Query(iconfilesSQL, iconId)
		if err != nil {
			return fmt.Errorf("error while retrieving iconfiles for '%s' from database: %w", iconName, err)
		}
		defer rows.Close()
		var format string
		var size string
		for rows.Next() {
			err = rows.Scan(&format, &size)
			if err != nil {
				return fmt.Errorf("error while retrieving iconfiles for '%s' from database: %w", iconName, err)
			}
			iconfiles = append(iconfiles, domain.IconfileDescriptor{
				Format: format,
				Size:   size,
			})
		}
		return nil
	}()
	if err != nil {
		return emptyIcon, err
	}

	tags := make([]string, 0, 50)
	err = func() error {
		rows, err = tx.Query(tagsSQL, iconId)
		if err != nil {
			return fmt.Errorf("error while retrieving tags for '%s' from database: %w", iconName, err)
		}
		defer rows.Close()
		var tag string
		for rows.Next() {
			err = rows.Scan(&tag)
			if err != nil {
				return fmt.Errorf("error while retrieving tags for '%s' from database: %w", iconName, err)
			}
			tags = append(tags, tag)
		}
		return nil
	}()
	if err != nil {
		return emptyIcon, err
	}

	return domain.IconDescriptor{
		IconAttributes: domain.IconAttributes{
			Name:       iconName,
			ModifiedBy: modifiedBy,
			Tags:       tags,
		},
		Iconfiles: iconfiles,
	}, nil
}

type PgRepository struct {
	logger zerolog.Logger
	Conn   connection
}

func NewPgRepository(conn connection) PgRepository {
	return PgRepository{
		Conn: conn,
	}
}

func (repo PgRepository) Close() error {
	return repo.Conn.Pool.Close()
}

// DescribeIcon returns the attributes of the icon having the specified name, "attributes" meaning here the entire icon without iconfiles' contents
func (repo PgRepository) DescribeIcon(ctx context.Context, iconName string) (domain.IconDescriptor, error) {
	tx, err := repo.Conn.Pool.Begin()
	if err != nil {
		return domain.IconDescriptor{}, err
	}
	defer tx.Rollback()
	return describeIconInTx(tx, iconName, false)
}

func (repo PgRepository) DescribeAllIcons(ctx context.Context) ([]domain.IconDescriptor, error) {
	tx, err := repo.Conn.Pool.Begin()
	if err != nil {
		return []domain.IconDescriptor{}, err
	}
	defer tx.Rollback()

	rows, errQuery := tx.Query("SELECT name FROM icon")
	if errQuery != nil {
		return []domain.IconDescriptor{}, fmt.Errorf("failed to retrieve all icon names: %w", errQuery)
	}
	defer rows.Close()

	iconNames := []string{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		iconNames = append(iconNames, name)
	}
	errProcessRows := rows.Err()
	if errProcessRows != nil {
		return []domain.IconDescriptor{}, fmt.Errorf("error while processing rows: %w", errProcessRows)
	}

	result := []domain.IconDescriptor{}
	for _, iconName := range iconNames {
		icon, errIconDesc := describeIconInTx(tx, iconName, false)
		if errIconDesc != nil {
			return []domain.IconDescriptor{}, fmt.Errorf("failed to retrieve icon %s: %w", iconName, errIconDesc)
		}
		result = append(result, icon)
	}

	return result, nil
}

func (repo PgRepository) CreateIcon(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor, modifiedBy string, createSideEffect func() error) error {
	var tx *sql.Tx
	var err error
	tx, err = repo.Conn.Pool.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction when creating icon %v: %w", iconName, err)
	}
	defer tx.Rollback()

	const insertIconSQL string = "INSERT INTO icon(name, modified_by) VALUES($1, $2) RETURNING id"
	_, err = tx.Exec(insertIconSQL, iconName, modifiedBy)
	if err != nil {
		reportErr := err
		if IsDBError(err, ErrDuplicateRows) {
			reportErr = domain.ErrIconAlreadyExists
		}
		return fmt.Errorf("failed to create icon %v: %w", iconName, reportErr)
	}

	err = insertIconfile(tx, iconName, iconfile)
	if err != nil {
		return fmt.Errorf("failed to create iconfile for %v: %w", iconName, err)
	}

	if createSideEffect != nil {
		err = createSideEffect()
		if err != nil {
			return fmt.Errorf("failed to create iconfile %s due to error while creating side-effect, %w", iconName, err)
		}
	}

	repo.logger.Info().Str("icon-name", iconName).Interface("iconfile", iconfile).Msg("Icon created")
	tx.Commit()
	return nil
}

func updateModifier(tx *sql.Tx, iconName string, modifiedBy string) error {
	_, err := tx.Exec("UPDATE icon SET modified_by = $1 WHERE name = $2", modifiedBy, iconName)
	if err != nil {
		return fmt.Errorf("failed to update icon %s with the modifier %s: %w", iconName, modifiedBy, err)
	}
	return nil
}

func (repo PgRepository) AddIconfileToIcon(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor, modifiedBy string, createSideEffect func() error) error {
	var tx *sql.Tx
	var err error

	tx, err = repo.Conn.Pool.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction when creating iconfile for %v: %w", iconName, err)
	}
	defer tx.Rollback()

	err = insertIconfile(tx, iconName, iconfile)
	if err != nil {
		return fmt.Errorf("failed to create iconfile %v: %w", iconName, err)
	}

	err = updateModifier(tx, iconName, modifiedBy)
	if err != nil {
		return fmt.Errorf("failed to add iconfile '%v' to icon '%s': %w", iconfile, iconName, err)
	}

	if createSideEffect != nil {
		err = createSideEffect()
		if err != nil {
			return fmt.Errorf("failed to create icon file %s due to error while creating side-effect: %w", iconName, err)
		}
	}

	tx.Commit()
	return nil
}

func insertIconfile(tx *sql.Tx, iconName string, iconfile domain.IconfileDescriptor) error {
	const insertIconfileSQL = "INSERT INTO icon_file(icon_id, file_format, icon_size) " +
		"SELECT id, $2, $3 FROM icon WHERE name = $1 RETURNING id"
	_, err := tx.Exec(insertIconfileSQL, iconName, iconfile.Format, iconfile.Size)
	if err != nil {
		if IsDBError(err, ErrDuplicateRows) {
			return domain.ErrIconfileAlreadyExists
		}
		return fmt.Errorf("failed to insert iconfile %v: %w", iconName, err)
	}
	return nil
}

func (repo PgRepository) GetExistingTags(tx context.Context) ([]string, error) {
	rows, err := repo.Conn.Pool.Query("SELECT text FROM tag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]string, 0, 50)
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve all tags: %w", err)
		}
		tags = append(tags, tag)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all tags: %w", err)
	}

	return tags, nil
}

func createTag(tx *sql.Tx, tag string) (int64, error) {
	var id int64
	// The lib/pq people messed the API up :-( : https://github.com/lib/pq/issues/24#issuecomment-841794798
	err := tx.QueryRow("INSERT INTO tag(text) VALUES($1) RETURNING id", tag).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last inserted id: %w", err)
	}
	return id, nil
}

func addTagReferenceToIcon(tx *sql.Tx, tagId int64, iconName string) error {
	_, err := tx.Exec("INSERT INTO icon_to_tags(icon_id, tag_id) SELECT id, $1 FROM icon WHERE name = $2", tagId, iconName)
	if err != nil {
		return fmt.Errorf("failed to add tag reference %d to icon '%s': %w", tagId, iconName, err)
	}
	return nil
}

func removeTagReferenceFromIcon(tx *sql.Tx, tagId int64, iconName string) error {
	_, err := tx.Exec(`
		DELETE FROM icon_to_tags
		WHERE tag_id = $1
			and icon_id = (SELECT id FROM icon WHERE name = $2)
	`, tagId, iconName)
	if err != nil {
		return fmt.Errorf("failed to remove tag reference %d from icon \"%s\": %w", tagId, iconName, err)
	}
	return nil
}

func GetTagId(tx *sql.Tx, tag string) (int64, error) {
	var tagId int64
	err := tx.QueryRow("SELECT id FROM tag WHERE text = $1", tag).Scan(&tagId)
	if err != nil {
		if err == sql.ErrNoRows {
			return createTag(tx, tag)
		}
		return 0, err
	}
	return tagId, nil
}

func (repo PgRepository) AddTag(ctx context.Context, iconName string, tag string, modifiedBy string) error {
	tx, trError := repo.Conn.Pool.Begin()
	if trError != nil {
		return fmt.Errorf("failed to obtain transaction for adding tag '%s' to '%s': %w", tag, iconName, trError)
	}
	defer tx.Rollback()

	tagId, insertTagErr := GetTagId(tx, tag)
	if insertTagErr != nil {
		return fmt.Errorf("failed to insert tag '%s' for '%s': %w", tag, iconName, insertTagErr)
	}
	addRefErr := addTagReferenceToIcon(tx, tagId, iconName)
	if addRefErr != nil {
		return fmt.Errorf("failed to connect tag '%s' to icon '%s': %w", tag, iconName, addRefErr)
	}

	err := updateModifier(tx, iconName, modifiedBy)
	if err != nil {
		return fmt.Errorf("failed to add tag '%s' to icon '%s': %w", tag, iconName, err)
	}

	tx.Commit()
	return nil
}

func (repo PgRepository) RemoveTag(ctx context.Context, iconName string, tag string, modifiedBy string) error {
	tx, trError := repo.Conn.Pool.Begin()
	if trError != nil {
		return fmt.Errorf("failed to obtain transaction for removing tag '%s' to '%s': %w", tag, iconName, trError)
	}
	defer tx.Rollback()

	tagId, insertTagErr := GetTagId(tx, tag)
	if insertTagErr != nil {
		return fmt.Errorf("failed to insert tag '%s' for '%s': %w", tag, iconName, insertTagErr)
	}
	removeRefErr := removeTagReferenceFromIcon(tx, tagId, iconName)
	if removeRefErr != nil {
		return fmt.Errorf("failed to disconnect tag '%s' from icon '%s': %w", tag, iconName, removeRefErr)
	}

	err := updateModifier(tx, iconName, modifiedBy)
	if err != nil {
		return fmt.Errorf("failed to remove tag '%s' from icon '%s': %w", tag, iconName, err)
	}

	tx.Commit()
	return nil
}

func deleteIconfileBare(tx *sql.Tx, iconName string, iconfile domain.IconfileDescriptor) (sql.Result, error) {
	var err error
	var sqlResult sql.Result

	var getIdAndLockIcon = "SELECT id FROM icon WHERE name = $1 FOR UPDATE"
	var deleteFile = "DELETE FROM icon_file WHERE icon_id = $1 and file_format = $2 and icon_size = $3"
	var countIconfilesLeftForIcon = "SELECT count(*) as icon_file_count FROM icon_file WHERE icon_id = $1"
	var deleteIconSQL = "DELETE FROM icon WHERE id = $1"

	var iconId int64
	err = tx.QueryRow(getIdAndLockIcon, iconName).Scan(&iconId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("icon %s not found: %w", iconName, domain.ErrIconNotFound)
		}
		return nil, fmt.Errorf("failed to obtain iconfile id for %v: %w", iconfile, err)
	}

	sqlResult, err = tx.Exec(deleteFile, iconId, iconfile.Format, iconfile.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to delete iconfile %v: %w", iconfile, err)
	}

	var remainingIconfileCountForIcon int
	err = tx.QueryRow(countIconfilesLeftForIcon, iconId).Scan(&remainingIconfileCountForIcon)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain iconfile count for %v: %w", iconName, err)
	}

	if remainingIconfileCountForIcon == 0 {
		_, err = tx.Exec(deleteIconSQL, iconId)
		if err != nil {
			return nil, fmt.Errorf("failed to delete icon %v: %w", iconName, err)
		}
	}

	return sqlResult, nil
}

func (repo PgRepository) DeleteIcon(ctx context.Context, iconName string, modifiedBy string, createSideEffect func() error) error {
	var tx *sql.Tx
	var err error

	tx, err = repo.Conn.Pool.Begin()
	if err != nil {
		return fmt.Errorf("failed to start Tx for deleting icon %s: %w", iconName, err)
	}
	defer tx.Rollback()

	var iconDesc domain.IconDescriptor
	iconDesc, err = describeIconInTx(tx, iconName, true)
	if err != nil {
		return fmt.Errorf("failed to describe icon %v: %w", iconName, err)
	}

	for _, iconFile := range iconDesc.Iconfiles {
		_, err = deleteIconfileBare(tx, iconName, iconFile)
		if err != nil {
			return fmt.Errorf("failed to delete iconfile %v: %w", iconFile, err)
		}
	}

	if createSideEffect != nil {
		err = createSideEffect()
		if err != nil {
			return fmt.Errorf("failed to execute side effect while deleting icon %v: %w", iconName, err)
		}
	}

	tx.Commit()
	return nil
}

func (repo PgRepository) DeleteIconfile(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor, modifiedBy string, createSideEffect func() error) error {
	var err error
	var tx *sql.Tx
	var sqlResult sql.Result
	var rowsAffected int64

	tx, err = repo.Conn.Pool.Begin()
	if err != nil {
		return fmt.Errorf("failed to create TX for deleting iconfile %v from %s: %w", iconfile, iconName, err)
	}
	defer tx.Rollback()

	sqlResult, err = deleteIconfileBare(tx, iconName, iconfile)
	if err != nil {
		return fmt.Errorf("failed to delete iconfile %v from %s: %w", iconfile, iconName, err)
	}
	rowsAffected, err = sqlResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected by deleting iconfile %v from %s: %w", iconfile, iconName, err)
	}
	if rowsAffected < 1 {
		return domain.ErrIconfileNotFound
	}

	err = updateModifier(tx, iconName, modifiedBy)
	if err != nil {
		return fmt.Errorf("failed to delete iconfile %v from icon '%s': %w", iconfile, iconName, err)
	}

	if createSideEffect != nil {
		err = createSideEffect()
		if err != nil {
			return fmt.Errorf("failed to create side-effect for removing iconfile %v from %s: %w", iconfile, iconName, err)
		}
	}

	tx.Commit()
	return nil
}
