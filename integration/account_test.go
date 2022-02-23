//nolint
package main

import (
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/gaikwadamolraj/form3"
	"github.com/gaikwadamolraj/form3/model"
)

var response model.AccountData
var errorMessage string

func GetAccount(accountId string) error {
	res, err := form3.FetchById(accountId)
	if err != nil {
		return fmt.Errorf("%s account not found", accountId)
	}

	response = res
	return nil
}

func GetAccountWithError(accountId string) error {
	_, err := form3.FetchById(accountId)
	if err != nil {
		errorMessage = err.Error()
		return nil
	}

	return nil
}

func accountConfirmStatus(accountId, status string) error {
	if response.ID == accountId && *response.Attributes.Status == status {
		return nil
	}
	return fmt.Errorf("accountId %s and status %s  not found", accountId, status)
}

func createAccountTest(accountId, status string) error {
	accountData := model.GetAccountModel()

	accountData.SetAccountID(accountId)
	accountData.SetStatus(status)

	_, err := form3.Create(accountData)
	if err != nil {
		return err
	}
	return nil
}

func deleteAccount(accountId string, version int, validation string) error {
	err := form3.DeleteByIdAndVer(accountId, version)
	if err != nil {
		if validation == "false" {
			return err
		}
		errorMessage = err.Error()
		return nil
	}

	return nil
}

func createAccountWithData(key, value string) error {
	accountData := model.GetAccountModel()

	if key == "accountId" {
		accountData.SetAccountID(value)
	} else if key == "status" {
		accountData.SetStatus(value)
	} else if key == "orgId" {
		accountData.SetOrgId(value)
	} else if key == "country" {
		accountData.SetCountry(value)
	}

	_, err := form3.Create(accountData)
	if err != nil {
		errorMessage = err.Error()
	}

	return nil
}

func validateError(err string) error {
	if false == strings.Contains(errorMessage, err) {
		return fmt.Errorf(fmt.Sprintf("actual: %s , expected : %s", err, errorMessage))
	}
	return nil
}

func validateGiven() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I created account (.*) and (.*)$`, createAccountTest)
	ctx.Step(`^Create account with (.*) and (.*) with validataion$`, createAccountWithData)

	ctx.Step(`^I get account (.*)$`, GetAccount)
	ctx.Step(`^I get error account (.*)$`, GetAccountWithError)
	ctx.Step(`^Validate account with id (.*) and status (.*)$`, accountConfirmStatus)
	ctx.Step(`^Delete the account (.*) and (.*) with error (.*)$`, deleteAccount)
	ctx.Step(`^Get error message (.*)$`, validateError)
	ctx.Step(`^Validate scenario$`, validateGiven)
}
