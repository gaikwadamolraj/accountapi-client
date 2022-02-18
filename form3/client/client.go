package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

func (c *Client) SendRequest(req *http.Request, mapRes interface{}) error {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	log.Printf("%d | %s - \"%s\" ", res.StatusCode, req.Method, req.URL)
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes model.ErrorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.ErrMessage)
		}
		return fmt.Errorf(fmt.Sprintf("Unknown error, status code: %d", res.StatusCode))
	}

	if mapRes != nil || mapRes == "" {
		response := model.Data{
			Data: mapRes,
		}

		return json.NewDecoder(res.Body).Decode(&response)
	}

	return nil
}
