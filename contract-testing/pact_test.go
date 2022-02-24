package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gaikwadamolraj/form3"
	"github.com/gaikwadamolraj/form3/model"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/suite"

	"net/http"
	"testing"
)

var ctx = context.Background()

type AccountClientSuite struct {
	suite.Suite
}

var (
	pact *dsl.Pact
)

func TestAccountClientSuite(t *testing.T) {
	suite.Run(t, new(AccountClientSuite))
}

func (s *AccountClientSuite) SetupSuite() {
	pact = &dsl.Pact{
		Consumer:                 "AccountApiClient",
		Provider:                 "AccountApi",
		PactDir:                  "./pacts",
		DisableToolValidityCheck: true,
		Host:                     "localhost",
	}
}

func (s *AccountClientSuite) TearDownSuite() {
	pact.Teardown()
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty" pact:"example=ad27e265-9605-4b4b-a0e5-3003ea9cc99a"`
	OrganisationID string             `json:"organisation_id,omitempty" pact:"example=ad27e265-9605-4b4b-a0e5-3003ea9cc99a"`
	Type           string             `json:"type,omitempty" pact:"example=accounts"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty" pact:"example=pending"`
	Switched                *bool    `json:"switched,omitempty"`
}

type Data struct {
	Data AccountData `json:"data,omitempty"`
}

func GetParsedErrorMsg(res *http.Response) string {
	var errorRes model.ErrorResponse
	_ = json.NewDecoder(res.Body).Decode(&errorRes)
	return errorRes.ErrMessage
}

func GetJsonRes(res *http.Response, mapRes *model.AccountData) {
	response := model.Data{
		Data: mapRes,
	}

	_ = json.NewDecoder(res.Body).Decode(&response)
}

func validateAssert(res *http.Response, expectedError string, expectedStatusCode int) error {
	actualErrorMsg := GetParsedErrorMsg(res)
	if actualErrorMsg != expectedError {
		return fmt.Errorf("Expected  \"%s\" but got %s and statusCode %d but got %d", expectedError, actualErrorMsg, expectedStatusCode, res.StatusCode)
	}

	return nil
}
func (s *AccountClientSuite) TestGetAccountThatDoesNotExist() {
	pact.AddInteraction().
		Given("Accound id  ad27e265-9605-4b4b-a0e5-3003ea9cc99c is not available").
		UponReceiving("A GET request for Account with id ad27e265-9605-4b4b-a0e5-3003ea9cc99c").
		WithRequest(
			dsl.Request{
				Method: http.MethodGet,
				Path:   dsl.String("/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc99c"),
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusNotFound,
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))

		res, _ := form3.FetchById(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc99c")

		if res.StatusCode != http.StatusNotFound {
			return fmt.Errorf("Expected 404 but got %d", res.StatusCode)
		}

		return nil
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestGetAccountThatExist() {
	pact.AddInteraction().
		Given("Accound id  ad27e265-9605-4b4b-a0e5-3003ea9cc99a is available").
		UponReceiving("A GET request for Account with id ad27e265-9605-4b4b-a0e5-3003ea9cc99a").
		WithRequest(
			dsl.Request{
				Method: http.MethodGet,
				Path:   dsl.String("/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc99a"),
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusOK,
				Body:   dsl.Match(Data{}),
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		response, _ := form3.FetchById(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc99a")

		s.Equal(http.StatusOK, response.StatusCode)

		return nil
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestCreateAccount() {
	pact.AddInteraction().
		Given("Create new account with id ad27e265-9605-4b4b-a0e5-3003ea9cc99a").
		UponReceiving("Create account for Account with id ad27e265-9605-4b4b-a0e5-3003ea9cc99a").
		WithRequest(
			dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String("/v1/organisation/accounts"),
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusCreated,
				Body:   dsl.Match(Data{}),
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		response, _ := form3.Create(ctx, model.GetAccountModel())
		acc := model.AccountData{}
		GetJsonRes(response, &acc)

		s.Equal("ad27e265-9605-4b4b-a0e5-3003ea9cc99a", acc.GetAccountID())
		s.Equal("pending", acc.GetStatus())

		return nil
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestCreateAccountIdValidation() {
	type ErrorResponse struct {
		ErrMessage string `json:"error_message,omitempty" pact:"example=id in body must be of type uuid"`
	}
	pact.AddInteraction().
		Given("Create new account with id 123").
		UponReceiving("Create account for Account with id 123").
		WithRequest(
			dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String("/v1/organisation/accounts"),
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusBadRequest,
				Body:   dsl.Match(ErrorResponse{}),
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		acc := model.GetAccountModel()
		acc.SetAccountID("123")
		res, _ := form3.Create(ctx, acc)

		return validateAssert(res, "id in body must be of type uuid", 400)
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestCreateAccountStatusValidation() {
	type ErrorResponse struct {
		ErrMessage string `json:"error_message,omitempty" pact:"example=status in body should be one of [pending confirmed failed]"`
	}
	pact.AddInteraction().
		Given("Create new account with status as Pending").
		UponReceiving("Create account for Account with status Pending").
		WithRequest(
			dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String("/v1/organisation/accounts"),
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusBadRequest,
				Body:   dsl.Match(ErrorResponse{}),
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		acc := model.GetAccountModel()
		acc.SetAccountID("123")
		res, _ := form3.Create(ctx, acc)

		return validateAssert(res, "status in body should be one of [pending confirmed failed]", 400)
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestCreateAccountDuplicateAccount() {
	type ErrorResponse struct {
		ErrMessage string `json:"error_message,omitempty" pact:"example=Account cannot be created as it violates a duplicate constraint"`
	}
	pact.AddInteraction().
		Given("Create new account two times with duplicate error").
		UponReceiving("Create account for two times with duplicate error").
		WithRequest(
			dsl.Request{
				Method: http.MethodPost,
				Path:   dsl.String("/v1/organisation/accounts"),
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusConflict,
				Body:   dsl.Match(ErrorResponse{}),
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		acc := model.GetAccountModel()
		res, _ := form3.Create(ctx, acc)

		return validateAssert(res, "Account cannot be created as it violates a duplicate constraint", 429)
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestDeleteAccountInvalidVersion() {
	type ErrorResponse struct {
		ErrMessage string `json:"error_message,omitempty" pact:"example=invalid version"`
	}
	pact.AddInteraction().
		Given("Delete account with inavalid version for id ad27e265-9605-4b4b-a0e5-3003ea9cc9ac").
		UponReceiving("Delete account with inavalid version for id ad27e265-9605-4b4b-a0e5-3003ea9cc9ac").
		WithRequest(
			dsl.Request{
				Method: http.MethodDelete,
				Path:   dsl.String("/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc9ac"),
				Query:  dsl.MapMatcher{"version": dsl.String("123")},
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusConflict,
				Body:   dsl.Match(ErrorResponse{}),
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))

		res, _ := form3.DeleteByIdAndVer(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc9ac", 123)

		return validateAssert(res, "invalid version", 429)
	}

	s.NoError(pact.Verify(test))
}

func (s *AccountClientSuite) TestDeleteAccountSuccess() {
	pact.AddInteraction().
		Given("Delete account with  id ad27e265-9605-4b4b-a0e5-3003ea9cc9mc").
		UponReceiving("Delete account with  id ad27e265-9605-4b4b-a0e5-3003ea9cc9mc").
		WithRequest(
			dsl.Request{
				Method: http.MethodDelete,
				Path:   dsl.String("/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc9mc"),
				Query:  dsl.MapMatcher{"version": dsl.String("123")},
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusNoContent,
			},
		)

	test := func() error {
		os.Setenv("API_HOST", fmt.Sprintf("http://localhost:%d", pact.Server.Port))

		res, _ := form3.DeleteByIdAndVer(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc9mc", 123)

		if res.StatusCode != http.StatusNoContent {
			return fmt.Errorf("Expected status code %d but got  %d", http.StatusNoContent, res.StatusCode)
		}

		return nil
	}

	s.NoError(pact.Verify(test))
}

// Merge pact with Bdd
