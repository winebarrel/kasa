package esa

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDriverDeleteOk(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://api.esa.io/v1/teams/example/posts/1",
		httpmock.NewStringResponder(http.StatusNoContent, ""))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	err := driver.Delete(1)
	assert.NoError(err)
}
