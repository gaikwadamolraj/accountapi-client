package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorStruct(t *testing.T) {
	errStruct := ErrorResponse{}

	assert.NotNil(t, errStruct, "Struct should be null")
	assert.Equal(t, errStruct.ErrMessage, "", "ErrorMessage value should be null")
}
