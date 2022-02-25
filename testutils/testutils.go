//nolint
package testutils

import (
	"context"
	"net/http"

	"github.com/gaikwadamolraj/form3"
	"github.com/gaikwadamolraj/form3/model"
)

func CreateAccount(acc *model.AccountData) (*http.Response, error) {
	ctx := context.Background()
	client := form3.NewClient()

	req, _ := client.Create(ctx, acc)
	return client.Do(req)
}

func DeleteRecord(accountId string, version int) (*http.Response, error) {
	ctx := context.Background()
	client := form3.NewClient()

	req, _ := client.Delete(ctx, accountId, version)
	return client.Do(req)
}

func FetRecord(accountId string) (*http.Response, error) {
	ctx := context.Background()
	client := form3.NewClient()

	req, _ := client.Fetch(ctx, accountId)
	return client.Do(req)
}
