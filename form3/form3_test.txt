package form3

import (
	"testing"

	"github.com/gaikwadamolraj/form3/model"
	"github.com/gaikwadamolraj/form3/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	//Create the accountObject
	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	accountData := model.GetAccountModel()
	accountData.SetAccountID("")
	accountData.SetOrgId(utils.GetUUID())
	accountData.SetCountry("GB")
	accountData.SetStatus("confirmed")

	// Invalid accountId
	accountData.SetAccountID("")
	_, err := Create(accountData)
	assert.NotNil(t, err, "Invalid accountId")
	assert.Containsf(t, err.Error(), "id in body is required", "Invalid accountId", "accountId")

	// Invalid account uuid
	accountData.SetAccountID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc123")
	_, err = Create(accountData)
	assert.NotNil(t, err, "Invalid account type uuid")
	assert.Containsf(t, err.Error(), "id in body must be of type uuid", "Invalid account type uuid", "accountId")

	// orgId is required
	accountData.SetAccountID(accountId)
	accountData.SetOrgId("")
	_, err = Create(accountData)
	assert.NotNil(t, err, "Invalid organisation_id")
	assert.Containsf(t, err.Error(), "organisation_id in body is required", "Invalid organisation_id", "OrgId")

	//Invalid orgId
	accountData.SetOrgId("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc123")
	_, err = Create(accountData)
	assert.NotNil(t, err, "organisation_id must be present")
	assert.Containsf(t, err.Error(), "organisation_id in body must be of type uuid", "organisation_id must be present", "OrgId")

	// country is required
	accountData.SetOrgId(utils.GetUUID())
	accountData.SetCountry("")
	_, err = Create(accountData)
	assert.NotNil(t, err, "country must be present")
	assert.Containsf(t, err.Error(), "country in body should match", "country must be present", "Country")

	//Invalid country with one numeric
	accountData.SetCountry("G1")
	_, err = Create(accountData)
	assert.NotNil(t, err, "Invalid country with one numeric")
	assert.Containsf(t, err.Error(), "country in body should match", "Invalid country with one numeric", "Country")

	//Invalid country with 3 char
	accountData.SetCountry("GBR")
	_, err = Create(accountData)
	assert.NotNil(t, err, "Invalid country with 3 char")
	assert.Containsf(t, err.Error(), "country in body should match", "Invalid country with 3 char", "Country")

	//staus required
	accountData.SetCountry("GB")
	accountData.SetStatus("")
	_, err = Create(accountData)
	assert.NotNil(t, err, "Status required")
	assert.Containsf(t, err.Error(), "status in body should be one of [pending confirmed failed]", "Status required", "Status")

	//Other than [pending confirmed failed]  status
	accountData.SetStatus("error")
	_, err = Create(accountData)
	assert.NotNil(t, err, "Other than [pending confirmed failed]  status")
	assert.Containsf(t, err.Error(), "status in body should be one of [pending confirmed failed]", "Other than [pending confirmed failed]  status", "Status")

	// Invalid payload pass
	_, err = Create("")
	assert.Equal(t, err.Error(), "json: cannot unmarshal string into Go struct field AccountCreation.data of type models.NewAccount", "Invalid payload pass")

	//Success Test
	deleteActiveRecord(accountId)
	accountData.SetStatus("pending")
	res, err := Create(accountData)
	assert.Nil(t, err, "Success test with accountId match")
	assert.Equal(t, accountId, res.GetAccountID(), "Success test with accountId match")
	assert.Equal(t, "pending", res.GetStatus(), "pending status")

	//Duplicate error
	_, err = Create(accountData)
	assert.Equal(t, err.Error(), "Account cannot be created as it violates a duplicate constraint", "Duplicate voilation")
}
func TestFetchById(t *testing.T) {
	// Success test
	accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	res, err := FetchById(accountId)
	assert.Nil(t, err, "Success - Get matching the record")
	assert.Equal(t, accountId, res.GetAccountID(), "Get matching the record")

	// Invalid uuid
	accountId = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc123"
	res, err = FetchById(accountId)
	assert.NotNil(t, err, "Not valid UUID error")
	assert.Equal(t, err.Error(), "id is not a valid uuid", "Not valid UUID error")

	//AccountId is mandatory
	accountId = ""
	res, err = FetchById(accountId)
	assert.NotNil(t, err, "AccountId not null validation")
	assert.Equal(t, err.Error(), "accountID is mandatory", "AccountId not null validation")
}

func deleteActiveRecord(accountId string) {
	_ = DeleteByIdAndVer(accountId, 0)
}
func TestDeleteByIdAndVer(t *testing.T) {
	// Record not found with id
	err := DeleteByIdAndVer("d4d999d8-7e0e-4475-829f-7f56bb5686ac", 0)
	assert.NotNil(t, err, "Expecting the 404")
	assert.Equal(t, err.Error(), "Unknown error, status code: 404", "If record not present")

	// Invalid version
	err = DeleteByIdAndVer("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 123)
	assert.NotNil(t, err, "Expecting invalid version")
	assert.Equal(t, err.Error(), "invalid version", "Invalid version")

	// Id is not valid uuid
	err = DeleteByIdAndVer("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc1234", 0)
	assert.NotNil(t, err, "Expecting Invalid uuid")
	assert.Equal(t, err.Error(), "id is not a valid uuid", "Invalid UUID")

	// Id is mandatory
	err = DeleteByIdAndVer("", 0)
	assert.NotNil(t, err, "Expecting accountID is mandatory")
	assert.Equal(t, err.Error(), "accountID is mandatory", "accountID is mandatory")

	// Successfully delete the record
	err = DeleteByIdAndVer("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
	assert.Nil(t, err, "Expecting account successfully deleted")
}
