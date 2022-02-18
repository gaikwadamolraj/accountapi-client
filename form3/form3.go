package form3

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gaikwadamolraj/form3/client"
	"github.com/gaikwadamolraj/form3/model"
	"github.com/gaikwadamolraj/form3/utils"
)

func Create(payload interface{}) (model.AccountData, error) {
	c := client.NewClient()
	ctx := context.Background()

	accountDeatils := model.AccountData{}
	url := utils.GetAccountUrl("", "")
	req, _ := c.NewRequestWithContext(ctx, http.MethodPost, url, payload)
	err := c.SendRequest(req, &accountDeatils)

	return accountDeatils, err
}

func FetchById(accountId string) (model.AccountData, error) {
	accountDeatils := model.AccountData{}

	if err := utils.ValidateaccountId(accountId); err != nil {
		return accountDeatils, err
	}

	c := client.NewClient()
	ctx := context.Background()
	url := utils.GetAccountUrl(accountId, "")
	req, _ := c.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	err := c.SendRequest(req, &accountDeatils)

	return accountDeatils, err
}

func DeleteByIdAndVer(accountId string, version int) error {
	if err := utils.ValidateaccountId(accountId); err != nil {
		return err
	}

	c := client.NewClient()
	ctx := context.Background()
	url := utils.GetAccountUrl(accountId, strconv.Itoa(version))
	req, _ := c.NewRequestWithContext(ctx, http.MethodDelete, url, nil)

	return c.SendRequest(req, nil)
}
