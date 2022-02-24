package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gaikwadamolraj/form3/model"
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

func (c *Client) NewRequestWithContext(ctx context.Context, httpMethod string, reqUrl string, payload interface{}) (*http.Request, error) {
	url, err := url.Parse(reqUrl)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if payload != nil {
		payload := model.Data{
			Data: payload,
		}

		_ = json.NewEncoder(buf).Encode(payload)
	}

	return http.NewRequestWithContext(ctx, httpMethod, url.String(), buf)
}
