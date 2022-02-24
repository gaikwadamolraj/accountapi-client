package form3

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gaikwadamolraj/form3/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateFailures(t *testing.T) {
	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	deleteActiveRecord(accountId)
	ctx := context.Background()
	var errorRes model.ErrorResponse

	//Create the accountObject
	accountData := model.GetAccountModel()

	// Invalid accountId
	accountData.SetAccountID("")
	res, _ := Create(ctx, accountData)

	_ = json.NewDecoder(res.Body).Decode(&errorRes)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Containsf(t, errorRes.ErrMessage, "id in body is required", "Invalid accountId", "accountId")

	//Success Test
	deleteActiveRecord(accountId)
	accountData.SetAccountID(accountId)
	res, _ = Create(ctx, accountData)
	assert.Equal(t, http.StatusCreated, res.StatusCode, "Success test with account creation")

	//Duplicate error
	accountData.SetAccountID(accountId)
	res, _ = Create(ctx, accountData)
	assert.Equal(t, http.StatusConflict, res.StatusCode, "Conflict for create same account")
}
func TestFetchById(t *testing.T) {
	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	ctx := context.Background()
	createNewRecord(accountId)
	// Success test
	res, _ := FetchById(ctx, accountId)
	assert.Equal(t, http.StatusOK, res.StatusCode, "Get matching the record with 200 OK")

	// Invalid uuid
	res, _ = FetchById(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc123")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Not valid UUID error")

	// //AccountId is mandatory
	_, err := FetchById(ctx, "")
	assert.Equal(t, err.Error(), "accountID is mandatory", "Account Id is required")
}

func deleteActiveRecord(accountId string) {
	ctx := context.Background()
	_, _ = DeleteByIdAndVer(ctx, accountId, 0)
}
func createNewRecord(accountId string) {
	ctx := context.Background()
	//Create the accountObject
	accountData := model.GetAccountModel()
	accountData.SetAccountID(accountId)
	_, _ = Create(ctx, accountData)
}

func TestDeleteByIdAndVer(t *testing.T) {
	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	ctx := context.Background()
	deleteActiveRecord(accountId)

	// Record not found with id
	res, _ := DeleteByIdAndVer(ctx, "d4d999d8-7e0e-4475-829f-7f56bb5686ac", 0)
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "If record not present")

	// Invalid version
	createNewRecord(accountId)
	res, _ = DeleteByIdAndVer(ctx, accountId, 123)
	assert.Equal(t, http.StatusConflict, res.StatusCode, "Invalid version")

	// Id is not valid uuid
	res, _ = DeleteByIdAndVer(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc1234", 0)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Invalid UUID")

	// Id is mandatory
	_, err := DeleteByIdAndVer(ctx, "", 0)
	assert.Equal(t, err.Error(), "accountID is mandatory", "accountID is mandatory")

	// Successfully delete the record
	res, _ = DeleteByIdAndVer(ctx, accountId, 0)
	assert.Equal(t, http.StatusNoContent, res.StatusCode, "Expecting account successfully deleted")
}
