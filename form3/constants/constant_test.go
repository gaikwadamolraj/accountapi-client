package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstats(t *testing.T) {
	assert.Equal(t, ApiVersion, "/v1", "ApiVersion")
	assert.Equal(t, OrgAccountRoute, "/organisation/accounts", "OrgAccountRoute")
}
