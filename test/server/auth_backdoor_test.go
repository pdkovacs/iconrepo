package server

import (
	"net/http"
	"testing"

	"iconrepo/internal/app/security/authn"
	"iconrepo/internal/app/security/authr"
	blobstore_tests "iconrepo/test/repositories/blobstore"
	"iconrepo/test/repositories/indexing"
	"iconrepo/test/testdata"

	"github.com/stretchr/testify/suite"
)

type authBackDoorTestSuite struct {
	ApiTestSuite
}

func TestAuthBackDoorTestSuite(t *testing.T) {
	suite.Run(
		t,
		&authBackDoorTestSuite{
			ApiTestSuite: apiTestSuites(
				"apitests_backdoor",
				[]blobstore_tests.TestBlobstoreController{blobstore_tests.DefaultBlobstoreController},
				[]indexing.IndexTestRepoController{*indexing.DefaultIndexTestRepoController()},
			)[0],
		},
	)
}

func (s *authBackDoorTestSuite) BeforeTest(suiteName string, testName string) {
	s.initTestCaseConfig(testName)
	if suiteName != "authBackDoorTestSuite" {
		return
	}
	switch testName {
	case "TestBackDoorMustntBeAvailableByDefault":
		{
		}
	case "TestBackDoorShouldBeAvailableWhenEnabled":
		{
			s.config.EnableBackdoors = true
		}
	case "TestBackDoorShouldAllowToSetPrivileges":
		{
			s.config.EnableBackdoors = true
		}
	default:
		{
			panic("Unexpected testName: " + testName)
		}
	}
	startErr := s.startApp(s.config)
	if startErr != nil {
		s.FailNow("", "%v", startErr)
	}
}

func (s *authBackDoorTestSuite) TestBackDoorMustntBeAvailableByDefault() {
	session := s.Client.mustLogin()
	resp, err := session.get(&testRequest{
		path: authenticationBackdoorPath,
		json: true,
		body: []authr.PermissionID{},
	})
	s.NoError(err)
	s.Equal(http.StatusNotFound, resp.statusCode)
}

func (s *authBackDoorTestSuite) TestBackDoorShouldBeAvailableWhenEnabled() {
	session := s.Client.mustLogin()
	resp, err := session.setAuthorization(
		[]authr.PermissionID{},
	)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.statusCode)
}

func (s *authBackDoorTestSuite) TestBackDoorShouldAllowToSetPrivileges() {
	requestedAuthorization := []authr.PermissionID{"galagonya", "ide-oda"}
	userID := authn.LocalDomain.CreateUserID(testdata.DefaultCredentials.Username)
	expectedUserInfo := authr.UserInfo{
		UserId:      userID,
		Permissions: requestedAuthorization,
		DisplayName: userID.String(),
	}

	session := s.Client.mustLogin()

	resp, err := session.setAuthorization(requestedAuthorization)
	s.NoError(err)
	s.Equal(200, resp.statusCode)

	resp, errUserInfo := session.get(&testRequest{
		path:          authenticationBackdoorPath,
		respBodyProto: &authr.UserInfo{},
	})
	s.NoError(errUserInfo)
	s.Equal(http.StatusOK, resp.statusCode)
	s.Equal(&expectedUserInfo, resp.body)
}
