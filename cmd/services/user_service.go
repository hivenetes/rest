package services

import (
	"rest/cmd/domain"
	"rest/cmd/utils"
)

// GetUser ...
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
