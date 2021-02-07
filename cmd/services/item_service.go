package services

import (
	"net/http"
	"rest/cmd/domain"
	"rest/cmd/utils"
)

type itemService struct {
}

// ItemService ...
var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
