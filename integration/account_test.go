//nolint
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"github.com/gaikwadamolraj/form3"
	"github.com/gaikwadamolraj/form3/model"
)

var response *http.Response
var errorMessage string
var ctx = context.Background()

//  Util functions
func GetParsedErrorMsg(res *http.Response) string {
	var errorRes model.ErrorResponse
	json.NewDecoder(res.Body).Decode(&errorRes)
	return errorRes.ErrMessage
}

func GetJsonRes(res *http.Response, mapRes *model.AccountData) {
	response := model.Data{
		Data: mapRes,
	}

	json.NewDecoder(res.Body).Decode(&response)
}

func GetAccount(accountId string) error {
	res, err := form3.FetchById(ctx, accountId)
	response = res
	if err != nil {
		return fmt.Errorf("%s account not found", accountId)
	}

	return nil
}

func GetAccountWithError(accountId string) error {
	res, err := form3.FetchById(ctx, accountId)
	response = res
	if err != nil {
		errorMessage = GetParsedErrorMsg(res)
		return nil
	}
	if res.StatusCode != http.StatusOK {
		errorMessage = GetParsedErrorMsg(res)
		return nil
	}

	return nil
}

func accountConfirmStatus(accountId, status string) error {
	acc := model.AccountData{}
	GetJsonRes(response, &acc)
	if acc.GetAccountID() == accountId && acc.GetStatus() == status {
		return nil
	}
	return fmt.Errorf("accountId %s and status %s  not found", accountId, status)
}

func createAccount(accountId, status string) error {
	accountData := model.GetAccountModel()

	accountData.SetAccountID(accountId)
	accountData.SetStatus(status)

	res, _ := form3.Create(ctx, accountData)
	response = res
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("Got error while creating account")
	}
	return nil
}

func deleteAccount(accountId string, version int, validation string) error {
	res, _ := form3.DeleteByIdAndVer(ctx, accountId, version)
	response = res
	if res.StatusCode != http.StatusNoContent {
		if validation == "false" {
			return fmt.Errorf(GetParsedErrorMsg(res))
		}
		errorMessage = GetParsedErrorMsg(res)
		return nil
	} else {
		errorMessage = GetParsedErrorMsg(res)
		return nil
	}
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

	res, _ := form3.Create(ctx, accountData)
	response = res
	if res.StatusCode != http.StatusCreated {
		errorMessage = GetParsedErrorMsg(res)
		return nil
	} else {
		errorMessage = GetParsedErrorMsg(res)
	}

	return nil
}

func validateError(err string, statuscode int) error {
	if err != "" {
		if false == strings.Contains(errorMessage, err) && response.StatusCode != statuscode {
			return fmt.Errorf(fmt.Sprintf("actual: %s , expected : %s", err+" "+strconv.Itoa(statuscode), errorMessage+" "+strconv.Itoa(response.StatusCode)))
		}
	} else {
		if response.StatusCode != statuscode {
			return fmt.Errorf(fmt.Sprintf("actual: %d , expected : %d", statuscode, response.StatusCode))
		}
	}

	return nil
}

func validateGiven() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I created account (.*) and (.*)$`, createAccount)
	ctx.Step(`^Create account with key (.*) and value (.*) with validataion$`, createAccountWithData)

	ctx.Step(`^I get account (.*)$`, GetAccount)
	ctx.Step(`^I get error account (.*)$`, GetAccountWithError)
	ctx.Step(`^Validate account with id (.*) and status (.*)$`, accountConfirmStatus)
	ctx.Step(`^Delete the account (.*) and (.*) with error (.*)$`, deleteAccount)
	ctx.Step(`^Get error message error (.*) and statuscode (.*)$`, validateError)
	ctx.Step(`^Validate scenario$`, validateGiven)
}
