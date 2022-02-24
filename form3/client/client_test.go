package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client, "Client is not null")
}

func TestNewRequestWithContextErros(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	_, err := client.NewRequestWithContext(ctx, http.MethodDelete, "/amol", nil)
	assert.Nil(t, err, "Error should be empty")
}

func TestNewRequestWithContextSuccess(t *testing.T) {
	client := NewClient()
	ctx := context.Background()

	req, err := client.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/api/amol", "")
	assert.Nil(t, err, "Sucess req payload")
	assert.Equal(t, req.Method, http.MethodGet, "Matching httpMethod")
}

func TestInvalidUrlParse(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	//Invalid url parse
	path := "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require"

	_, err := c.NewRequestWithContext(ctx, http.MethodGet, path, nil)

	assert.NotNil(t, err, "Url parse error")
}
