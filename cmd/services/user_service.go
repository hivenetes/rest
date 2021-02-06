package services

import (
	"rest/cmd/utils"
	"rest/cmd/domain"
)

// GetUser ...
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
