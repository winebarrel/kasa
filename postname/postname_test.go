package postname

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostnameJoin(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string
		category string
		name     string
	}{
		{"foo/bar/zoo", "foo/bar", "zoo"},
		{"foo/bar", "foo", "bar"},
		{"foo/bar/", "foo/bar/", ""},
		{"foo", "", "foo"},
		{"", "", ""},
	}

	for _, t := range tests {
		cat, name := Split(t.input)
		assert.Equal(t.category, cat)
		assert.Equal(t.name, name)
	}
}
