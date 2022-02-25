package form3

import (
	"context"
	"net/http"
	"testing"

	"github.com/gaikwadamolraj/form3/model"
	"github.com/stretchr/testify/assert"
)

func createAccount(acc *model.AccountData) (*http.Response, error) {
	ctx := context.Background()
	client := NewClient()

	req, _ := client.Create(ctx, acc)
	return client.Do(req)
}

func deleteRecord(accountId string, version int) (*http.Response, error) {
	ctx := context.Background()
	client := NewClient()

	req, _ := client.Delete(ctx, accountId, version)
	return client.Do(req)
}

func fetRecord(accountId string) (*http.Response, error) {
	ctx := context.Background()
	client := NewClient()

	req, _ := client.Fetch(ctx, accountId)
	return client.Do(req)
}

func TestCreate(t *testing.T) {
	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	_, _ = deleteRecord(accountId, 0)

	//Create the accountObject
	accountData := model.GetAccountModel()

	//Success Test
	_, _ = deleteRecord(accountId, 0)
	accountData.SetAccountID(accountId)
	res, _ := createAccount(accountData)
	assert.Equal(t, http.StatusCreated, res.StatusCode, "Success test with account creation")
}
func TestFetch(t *testing.T) {
	client := NewClient()

	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	acc := model.GetAccountModel()
	_, _ = createAccount(acc)

	// Success test
	res, _ := fetRecord(accountId)
	assert.Equal(t, http.StatusOK, res.StatusCode, "Get matching the record with 200 OK")

	//AccountId is mandatory
	_, err := client.Fetch(context.Background(), "")
	assert.Equal(t, err.Error(), "accountID is mandatory", "Account Id is required")
}

func TestDelete(t *testing.T) {
	client := NewClient()

	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	// Successfully delete the record
	res, _ := deleteRecord(accountId, 0)
	assert.Equal(t, http.StatusNoContent, res.StatusCode, "Expecting account successfully deleted")

	//AccountId is mandatory
	_, err := client.Delete(context.Background(), "", 0)
	assert.Equal(t, err.Error(), "accountID is mandatory", "Account Id is required")
}
