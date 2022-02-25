package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gaikwadamolraj/form3/model"
	"github.com/gaikwadamolraj/form3/utils"
)

type Client struct {
	HttpClient *http.Client
}

func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: time.Minute,
	}

	return &Client{HttpClient: httpClient}
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	return client.HttpClient.Do(req)
}

func (client *Client) Create(ctx context.Context, payload *model.AccountData) (*http.Request, error) {
	url := utils.GetAccountUrl("", "")

	buf := &bytes.Buffer{}
	if payload != nil {
		payload := model.Data{
			Data: payload,
		}

		_ = json.NewEncoder(buf).Encode(payload)
	}

	return http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
}

func (client *Client) Fetch(ctx context.Context, accountId string) (*http.Request, error) {
	if err := utils.ValidateaccountId(accountId); err != nil {
		return nil, err
	}

	url := utils.GetAccountUrl(accountId, "")
	return http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
}

func (client *Client) Delete(ctx context.Context, accountId string, version int) (*http.Request, error) {
	if err := utils.ValidateaccountId(accountId); err != nil {
		return nil, err
	}
	url := utils.GetAccountUrl(accountId, strconv.Itoa(version))
	return http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
}
