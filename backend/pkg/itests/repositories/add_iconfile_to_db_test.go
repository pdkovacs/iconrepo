package repositories

import (
	"errors"
	"testing"

	"github.com/pdkovacs/igo-repo/backend/pkg/domain"
	"github.com/pdkovacs/igo-repo/backend/pkg/itests"
	"github.com/stretchr/testify/suite"
)

type addIconfileToDBTestSuite struct {
	dbTestSuite
}

func TestAddIconfileToDBTestSuite(t *testing.T) {
	suite.Run(t, &addIconfileToDBTestSuite{})
}

func (s *addIconfileToDBTestSuite) TestErrorOnDuplicateIconfile() {
	var err error
	var icon = itests.TestData[0]
	var iconfile = icon.Iconfiles[0]

	err = s.CreateIcon(icon.Name, iconfile, icon.ModifiedBy, nil)
	s.NoError(err)

	err = s.AddIconfileToIcon(icon.Name, iconfile, icon.ModifiedBy, nil)
	s.True(errors.Is(err, domain.ErrIconfileAlreadyExists))
}

func (s *addIconfileToDBTestSuite) TestSecondIconfile() {
	var err error
	var icon = itests.TestData[0]
	var iconfile1 = icon.Iconfiles[0]
	var iconfile2 = icon.Iconfiles[1]

	err = s.CreateIcon(icon.Name, iconfile1, icon.ModifiedBy, nil)
	s.NoError(err)

	err = s.AddIconfileToIcon(icon.Name, iconfile2, icon.ModifiedBy, nil)
	s.NoError(err)

	var iconDesc domain.Icon
	iconDesc, err = s.DescribeIcon(icon.Name)
	s.NoError(err)
	s.equalIconAttributes(icon, iconDesc, nil)
}

func (s *addIconfileToDBTestSuite) TestAddSecondIconfileBySecondUser() {
	var err error
	var icon = itests.TestData[0]
	var iconfile1 = icon.Iconfiles[0]
	var iconfile2 = icon.Iconfiles[1]

	var secondUser = "sedat"

	err = s.CreateIcon(icon.Name, iconfile1, icon.ModifiedBy, nil)
	s.NoError(err)

	err = s.AddIconfileToIcon(icon.Name, iconfile2, secondUser, nil)
	s.NoError(err)

	var iconDesc domain.Icon
	iconDesc, err = s.DescribeIcon(icon.Name)
	s.NoError(err)
	clone := itests.CloneIcon(icon)
	clone.ModifiedBy = secondUser
	s.equalIconAttributes(clone, iconDesc, nil)
}
