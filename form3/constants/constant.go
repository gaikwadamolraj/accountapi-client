package constants

import (
	"os"
)

const (
	ApiVersion      = "/v1"
	OrgAccountRoute = "/organisation/accounts"
)

var BaseUrl = os.Getenv("API_HOST")
