package services

import (
	"rest/cmd/domain"
	"rest/cmd/utils"
)

type userService struct {
}

// UserService ...
var (
	UserService userService
)

// GetUser ...
func (u *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDatabase.GetUser(userID)
}
