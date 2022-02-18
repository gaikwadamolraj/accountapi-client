package utils

import (
	"fmt"

	"github.com/gaikwadamolraj/form3/constants"
	"github.com/google/uuid"
)

func GetUUID() string {
	return uuid.New().String()
}

func GetAccountUrl(accountId string, version string) string {
	path := constants.BaseUrl + constants.ApiVersion + constants.OrgAccountRoute
	if version != "" && accountId != "" {
		path += "/" + accountId + fmt.Sprintf("?version=%s", version)
	} else if accountId != "" {
		path += "/" + accountId
	}

	return path
}

func ValidateaccountId(accountId string) error {
	if accountId == "" || len(accountId) == 0 {
		return fmt.Errorf("accountID is mandatory")
	}
	return nil
}
