package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/utils"
)

func TestTagsToString(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := []struct {
		tags     []string
		expected string
	}{
		{[]string{"a", "b", "c", "d"}, "[#a,#b,#c,#d]"},
		{[]string{}, ""},
	}

	for _, t := range tests {
		assert.Equal(utils.TagsToString(t.tags), t.expected)
	}
}
