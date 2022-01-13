package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPostNum(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name       string
		searchPath string
		expect     int
	}{
		{name: "esa directory path", searchPath: "foo/bar/zoo", expect: 0},
		{name: "url path", searchPath: "https://docs.esa.io/posts/1", expect: 1},
		{name: "post num", searchPath: "//1", expect: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := GetPostNum("docs", tt.searchPath)
			assert.NoError(err)
			assert.Equal(tt.expect, n)
		})
	}
}
