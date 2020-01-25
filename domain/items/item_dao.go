package items

import "github.com/waytkheming/bookstore-items-api/client/elasticsearch"

import "github.com/waytkheming/bookstore-items-api/utils/errors"

const (
	indexItems = "items"
)

func (i *Item) Save() *errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item")
	}
	i.ID = result.Id
	return nil
}
