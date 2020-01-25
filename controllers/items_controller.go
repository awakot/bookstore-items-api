package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/waytkheming/bookstore-items-api/domain/items"
	"github.com/waytkheming/bookstore-items-api/services"
	"github.com/waytkheming/bookstore-items-api/utils/errors"
	"github.com/waytkheming/bookstore-items-api/utils/formatter"
	"github.com/waytkheming/sample-oauth-golang/oauth"
)

var (
	ItemsController itemsControllerinterface = &itemsController{}
)

type itemsControllerinterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		formatter.RespondJSON(w, err.Code, err)
		return
	}
	authorID := oauth.GetCallerID(r)

	if authorID == 0 {
		restErr := errors.NewUnauthorizedError()
		formatter.RespondErr(w, *restErr)
		return
	}

	var itemReq items.Item
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid req body")
		formatter.RespondErr(w, *restErr)
		return
	}

	defer r.Body.Close()
	if err := json.Unmarshal(body, &itemReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		formatter.RespondErr(w, *restErr)
		return
	}

	itemReq.Author = authorID

	result, restErr := services.ItemService.Create(itemReq)
	if err != nil {
		formatter.RespondErr(w, *restErr)
		return
	}

	formatter.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
