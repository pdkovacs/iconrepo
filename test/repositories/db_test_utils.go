package repositories_tests

import (
	"database/sql"
	"fmt"

	"igo-repo/internal/app/domain"
	"igo-repo/internal/config"
	"igo-repo/internal/logging"
	"igo-repo/internal/repositories"
	common_test "igo-repo/test/common"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type DBTestSuite struct {
	suite.Suite
	config config.Options
	dbRepo *repositories.DBRepository
	logger zerolog.Logger
}

func DeleteDBData(db *sql.DB) error {
	var tx *sql.Tx
	var err error

	tx, err = db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start Tx for deleting test data: %w", err)
	}
	defer tx.Rollback()

	tables := []string{"icon", "icon_file", "tag", "icon_to_tags"}
	for _, table := range tables {
		_, err = tx.Exec("DELETE FROM " + table)
		if err != nil {
			if repositories.IsDBError(err, repositories.ErrMissingDBTable) {
				continue
			}
			return fmt.Errorf("failed to delete test data from table %s: %w", table, err)
		}
	}

	_ = tx.Commit()
	return nil
}

func (s *DBTestSuite) NewTestDBRepo() {
	var err error
	config := common_test.GetTestConfig()
	connection, err := repositories.NewDBConnection(config, logging.CreateUnitLogger(s.logger, "test-db-connection"))
	if err != nil {
		panic(fmt.Sprintf("failed to create test connection: %v", err))
	}
	_, schemaErr := repositories.OpenDBSchema(config, connection, logging.CreateUnitLogger(s.logger, "db-schema"))
	if schemaErr != nil {
		panic(schemaErr)
	}
	err = DeleteDBData(connection.Pool)
	if err != nil {
		panic(fmt.Sprintf("failed to delete test data: %v", err))
	}
	s.dbRepo = repositories.NewDBRepository(connection, logging.CreateUnitLogger(s.logger, "test-db-repo"))
	if err != nil {
		panic(err)
	}
}

func manageTestResourcesAfterEach() {
}

func (s *DBTestSuite) SetupSuite() {
	s.config.DBSchemaName = "itest_repositories"
	s.logger = logging.CreateRootLogger(logging.DebugLevel)
}

func (s *DBTestSuite) TearDownSuite() {
	s.dbRepo.Conn.Pool.Close()
}

func (s *DBTestSuite) BeforeTest(suiteName, testName string) {
	s.NewTestDBRepo()
}

func (s *DBTestSuite) AfterTest(suiteName, testName string) {
	manageTestResourcesAfterEach()
}

func (s *DBTestSuite) equalIconAttributes(icon1 domain.Icon, icon2 domain.IconDescriptor, expectedTags []string) {
	s.Equal(icon1.Name, icon2.Name)
	s.Equal(icon1.ModifiedBy, icon2.ModifiedBy)
	if expectedTags != nil {
		s.Equal(expectedTags, icon2.Tags)
	}
}

func (s DBTestSuite) getIconCount() (int, error) {
	var getIconCountSQL = "SELECT count(*) from icon"
	var count int
	err := s.dbRepo.Conn.Pool.QueryRow(getIconCountSQL).Scan(&count)
	if err != nil {
		return 0, nil
	}
	return count, nil
}

func (s *DBTestSuite) getIconfile(iconName string, iconfile domain.Iconfile) ([]byte, error) {
	return s.dbRepo.GetIconfile(iconName, iconfile.IconfileDescriptor)
}

func (s *DBTestSuite) getIconfileChecked(iconName string, iconfile domain.Iconfile) {
	content, err := s.getIconfile(iconName, iconfile)
	s.NoError(err)
	s.Equal(iconfile.Content, content)
}
