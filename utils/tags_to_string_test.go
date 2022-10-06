package utils_test

import (
	"testing"

	"github.com/kanmu/kasa/utils"
	"github.com/stretchr/testify/assert"
)

func TestTagsToString(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := []struct {
		tags     []string
		expected string
	}{
		{[]string{"a", "b", "c", "d"}, "[#a,#b,#c,#d]"},
		{[]string{"a", "#b", "##c", "###d"}, "[#a,#b,#c,#d]"},
		{[]string{}, ""},
	}

	for _, t := range tests {
		assert.Equal(t.expected, utils.TagsToString(t.tags))
	}
}
