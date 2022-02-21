package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gaikwadamolraj/form3"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/suite"

	"net/http"
	"testing"
)

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
		_, err := form3.FetchById("ad27e265-9605-4b4b-a0e5-3003ea9cc99c")
		log.Println("Here result ", err)
		if err.Error() != "Unknown error, status code: 404" {
			return errors.New("error returned not as expected")
		}

		return nil
	}

	s.NoError(pact.Verify(test))
}

// Only one test added right now.
// Will merge BDD (cucumber) testing with pact.
