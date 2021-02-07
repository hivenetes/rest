package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"rest/cmd/domain"
	"rest/cmd/utils"
	"testing"
)

type userDatabaseMock struct {
}

var (
	UserDatabaseMock userDatabaseMock
	getUserFunction  func(userID int64) (*domain.User, *utils.ApplicationError)
)


func (m *userDatabaseMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func init() {
	domain.UserDatabase = &UserDatabaseMock
}

func TestGetUserNotInDatabase(t *testing.T) {

	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 not found",
			StatusCode: http.StatusNotFound,
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "user 0 not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserInDatabase(t *testing.T) {

	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123}, nil
	}

	user, err := UserService.GetUser(123)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.ID)
}
