package feature

import (
	"testing"
	"encoding/json"
	"bytes"

	"github.com/stretchr/testify/suite"
	"karuhundeveloper.com/gostarterkit/app/models/user"
	"karuhundeveloper.com/gostarterkit/tests"
	"github.com/goravel/framework/facades"
)

type AuthTestSuite struct {
	suite.Suite
	tests.TestCase
	user *user.User
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) SetupTest() {
	password, err := facades.Hash().Make("password")
	s.Nil(err)

	s.user = &user.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: password,
	}
	s.Nil(facades.Orm().Query().Create(s.user))
}

func (s *AuthTestSuite) TearDownTest() {
	facades.Orm().Query().Delete(s.user)
}

func (s *AuthTestSuite) TestLogin() {
	// Test success
	payload := map[string]any{
		"email":    "test@example.com",
		"password": "password",
	}
	body, err := json.Marshal(payload)
	s.Nil(err)

	response, err := s.Http(s.T()).Post("/api/v1/auth/login", bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(200)
	
	// Test invalid credentials
	payload = map[string]any{
		"email":    "test@example.com",
		"password": "wrong_password",
	}
	body, err = json.Marshal(payload)
	s.Nil(err)

	response, err = s.Http(s.T()).Post("/api/v1/auth/login", bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(401)
}

func (s *AuthTestSuite) login() string {
	payload := map[string]any{
		"email":    "test@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(payload)
	response, err := s.Http(s.T()).Post("/api/v1/auth/login", bytes.NewBuffer(body))
	s.Nil(err)
	response.AssertStatus(200)

	data, err := response.Json()
	s.Nil(err)
	return data["data"].(map[string]any)["token"].(string)
}

func (s *AuthTestSuite) TestRefreshToken() {
	token := s.login()

	// Test success
	response, err := s.Http(s.T()).WithHeader("Authorization", "Bearer " + token).Put("/api/v1/auth/refresh-token", nil)
	s.Nil(err)
	response.AssertStatus(200)
}

func (s *AuthTestSuite) TestLogout() {
	token := s.login()

	// Test success
	response, err := s.Http(s.T()).WithHeader("Authorization", "Bearer " + token).Post("/api/v1/auth/logout", nil)
	s.Nil(err)
	response.AssertStatus(200)
}
