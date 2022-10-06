package utils_test

import (
	"testing"

	"github.com/kanmu/kasa/utils"
	"github.com/stretchr/testify/assert"
)

func TestTagContains(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := []struct {
		tags     []string
		tag      string
		expected bool
	}{
		{[]string{"a", "b", "c", "d"}, "a", true},
		{[]string{"#a", "#b", "c", "d"}, "#b", true},
		{[]string{"a", "b", "c", "d"}, "#a", true},
		{[]string{}, "b", false},
		{[]string{"a", "b", "c", "d"}, "e", false},
		{[]string{"#a", "#b", "c", "d"}, "#e", false},
	}

	for _, t := range tests {
		assert.Equal(t.expected, utils.TagContains(t.tags, t.tag))
	}
}
