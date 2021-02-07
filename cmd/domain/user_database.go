package domain

import (
	"fmt"
	"net/http"
	"rest/cmd/utils"
)

type userDatabaseInterface interface {
	GetUser(userID int64) (*User, *utils.ApplicationError)
}

type userDatabase struct{}

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Ab", LastName: "S", Email: "abs@mail.com"},
	}

	UserDatabase userDatabaseInterface
)

func init() {
	UserDatabase = &userDatabase{}
}

// GetUser ...
func (d *userDatabase) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userID),
		StatusCode: http.StatusNotFound,
	}
}
