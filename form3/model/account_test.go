package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountDataStrcut(t *testing.T) {
	accountData := AccountData{}
	assert.NotNil(t, accountData, "AccountData Struct should not be null")

}

func TestAccountAttributesStrcut(t *testing.T) {
	accAttrib := AccountAttributes{}
	assert.NotNil(t, accAttrib, "AccountAttributes Struct should not be null")
}

func TestDataStrcut(t *testing.T) {
	data := Data{}
	assert.NotNil(t, data, "Data Struct should not be null")
}

func TestAccountsStrcut(t *testing.T) {
	data := Accounts{}

	assert.NotNil(t, data, "Accounts Struct should not be null")
}

func TestGetAccountModel(t *testing.T) {
	acc := GetAccountModel()
	assert.NotNil(t, acc, "Accounts Struct should not be null")
	assert.Equal(t, acc.OrganisationID, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", "Org Id value check")
}

func TestGetint64Pointer(t *testing.T) {
	s := getint64Pointer(0)
	assert.Equal(t, *s, int64(0), "Get int64 pointer value and match again")
}

func TestSetStringPointer(t *testing.T) {
	s := setStringPointer("Amol")
	assert.Equal(t, *s, string("Amol"), "Get string pointer value and match again")
}

func TestAccountGetSet(t *testing.T) {
	acc := GetAccountModel()
	acc.SetAccountID("Amol")
	assert.Equal(t, acc.GetAccountID(), "Amol", "AccountId GetSet")
	acc.SetOrgId("Amol")
	assert.Equal(t, acc.GetOrgId(), "Amol", "OrgId GetSet")
	acc.SetType("Amol")
	assert.Equal(t, acc.GetType(), "Amol", "Type GetSet")
	assert.Equal(t, acc.GetVersion(), int64(-1), "Version GetSet - Negative")
	acc.SetVersion(0)
	assert.Equal(t, acc.GetVersion(), int64(0), "Version GetSet")

	acc.SetStatus(" ")
	assert.Equal(t, acc.GetStatus(), " ", "Status GetSet - Null")
	acc.SetStatus("Amol")
	assert.Equal(t, acc.GetStatus(), "Amol", "Status GetSet")

	acc.SetName("Amol")
	assert.Equal(t, acc.Attributes.Name[0], "Amol", "Name GetSet")

	acc.SetCountry("GB")
	acc.SetBaseCurrency("GBR")
	acc.SetBankID("123")
	acc.SetBankIDCode("code")
	acc.SetBic("Bic")
	acc.SetIban("iban")
}
