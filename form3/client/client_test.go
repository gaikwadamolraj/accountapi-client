package client

import (
	"context"
	"fmt"
	"github.com/gaikwadamolraj/form3/constants"
	"github.com/gaikwadamolraj/form3/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client, "Client is not null")
}

func TestNewRequestWithContext(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	_, err := client.NewRequestWithContext(ctx, http.MethodDelete, "/amol", nil)
	assert.Nil(t, err, "Error should be empty")

	_, err = client.NewRequestWithContext(ctx, "Delete", "amol", nil)
	assert.Nil(t, err, "Invalid url error")

	_, err = client.NewRequestWithContext(ctx, "", "/amol", nil)
	assert.Nil(t, err, "Error should be empty for blank Method")

	_, err = client.NewRequestWithContext(ctx, "", "/amol", "")
	assert.Nil(t, err, "Error should be empty as blank data will be passed")

	_, err = client.NewRequestWithContext(ctx, "", "/amol", nil)
	assert.Nil(t, err, "Error should be empty as blank data will be passed")

	payload := model.ErrorResponse{}
	_, err = client.NewRequestWithContext(ctx, "", "/amol", payload)
	assert.Nil(t, err, "Error should be empty as blank data will be passed")

	_, err = client.NewRequestWithContext(ctx, "", "/amol", "amolgaikwad")
	assert.Nil(t, err, "Error should be empty as blank data will be passed")
}

func TestSendRequest(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	// Invalid uuid passed
	path := constants.BaseUrl + constants.ApiVersion + constants.OrgAccountRoute + "/123" + fmt.Sprintf("?version=%d", 0)
	req, _ := c.NewRequestWithContext(ctx, http.MethodDelete, path, nil)
	err := c.SendRequest(req, "")
	assert.NotNil(t, err, "Delete Invalid uuid passed")

	//404
	path = constants.BaseUrl + constants.ApiVersion + constants.OrgAccountRoute + "/ad27e265-9605-4b4b-a0e5-3003ea9cc4dd" + fmt.Sprintf("?version=%d", 0)
	req, _ = c.NewRequestWithContext(ctx, http.MethodDelete, path, nil)
	err = c.SendRequest(req, "")
	assert.NotNil(t, err, "404 not found")

	//Invalid request
	path = constants.BaseUrl + constants.ApiVersion + constants.OrgAccountRoute + "/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	req, _ = c.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	err = c.SendRequest(req, "")
	assert.NotNil(t, err, "404 not found")

	//Req failed request
	path = "http://localhost:9" + constants.ApiVersion + constants.OrgAccountRoute + "/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	req, _ = c.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	err = c.SendRequest(req, "")
	assert.NotNil(t, err, "Connection refused")

	// POST sucesss
	acc := model.GetAccountModel()
	acc.SetAccountID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	acc.SetOrgId("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	acc.SetCountry("GB")
	mapRes := model.AccountData{}
	path = constants.BaseUrl + constants.ApiVersion + constants.OrgAccountRoute
	req, _ = c.NewRequestWithContext(ctx, http.MethodPost, path, acc)
	err = c.SendRequest(req, &mapRes)

	assert.Nil(t, err, "Successful create acc")

	// Delete sucess
	path = constants.BaseUrl + constants.ApiVersion + constants.OrgAccountRoute + "/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc" + fmt.Sprintf("?version=%d", 0)
	req, _ = c.NewRequestWithContext(ctx, http.MethodDelete, path, nil)
	err = c.SendRequest(req, nil)
	assert.Nil(t, err, "Delete record successfully")
}

func TestInvalidUrlParse(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	//Invalid url parse
	path := "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require"

	_, err := c.NewRequestWithContext(ctx, http.MethodGet, path, nil)

	assert.NotNil(t, err, "Url parse error")
}
