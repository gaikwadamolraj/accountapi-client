package form3

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gaikwadamolraj/form3/client"
	"github.com/gaikwadamolraj/form3/model"
	"github.com/gaikwadamolraj/form3/utils"
)

func Create(ctx context.Context, payload *model.AccountData) (*http.Response, error) {
	c := client.NewClient()

	url := utils.GetAccountUrl("", "")
	req, _ := c.NewRequestWithContext(ctx, http.MethodPost, url, payload)

	return c.HttpClient.Do(req)
}

func FetchById(ctx context.Context, accountId string) (*http.Response, error) {
	if err := utils.ValidateaccountId(accountId); err != nil {
		return nil, err
	}

	c := client.NewClient()
	url := utils.GetAccountUrl(accountId, "")
	req, _ := c.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	return c.HttpClient.Do(req)
}

func DeleteByIdAndVer(ctx context.Context, accountId string, version int) (*http.Response, error) {
	if err := utils.ValidateaccountId(accountId); err != nil {
		return nil, err
	}

	c := client.NewClient()
	url := utils.GetAccountUrl(accountId, strconv.Itoa(version))
	req, _ := c.NewRequestWithContext(ctx, http.MethodDelete, url, nil)

	return c.HttpClient.Do(req)
}
