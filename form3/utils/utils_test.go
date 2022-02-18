package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetUUID(t *testing.T) {
	id := GetUUID()
	splitArr := strings.Split(id, "-")
	assert.NotNil(t, id, "UUID is not null")
	assert.Equal(t, len(strings.Split(splitArr[0], "")), 8, "Fist portion should 8 chars")
	assert.Equal(t, len(splitArr), 5, "Length shoud be 5")
}

func TestValidateaccountId(t *testing.T) {
	err := ValidateaccountId(GetUUID())
	assert.Nil(t, err, "Valid UUId")

	err = ValidateaccountId("")
	assert.NotNil(t, err, "InValid accountId")
	assert.Equal(t, err.Error(), "accountID is mandatory", "Invalid account Id")
}

func TestGetAccountUrl(t *testing.T) {
	// Valid account org url
	url := GetAccountUrl("", "")
	assert.NotNil(t, url, "Valid url ")

	// valid url for url path id
	url = GetAccountUrl("123-123-123-123", "")
	assert.NotNil(t, url, "Valid url for accountId")
	assert.Containsf(t, url, "123-123-123-123", "Contains the accountid in url", "URL")

	// accountId with version
	url = GetAccountUrl("123-123-123-123", "123")
	assert.NotNil(t, url, "Valid url for accountId and version")
	assert.Containsf(t, url, "123-123-123-123", "Contains the accountid and version in url", "URL")
	assert.Containsf(t, url, "123", "Contains the accountid and version in url", "URL")
}
