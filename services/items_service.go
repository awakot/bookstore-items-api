package services

import "github.com/waytkheming/bookstore-items-api/domain/items"

import "github.com/waytkheming/bookstore-items-api/utils/errors"

var (
	ItemService itemsServiceInterface = &itemService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, *errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(ID string) (*items.Item, *errors.RestErr) {
	return nil, errors.NewNotFoundError("implement me")
}
