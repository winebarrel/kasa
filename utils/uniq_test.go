package utils_test

import (
	"testing"

	"github.com/kanmu/kasa/utils"
	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := []struct {
		ary      []string
		expected []string
	}{
		{[]string{"a", "b", "c", "a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"d", "c", "b", "a", "c", "b", "a"}, []string{"a", "b", "c", "d"}},
	}

	for _, t := range tests {
		assert.Equal(utils.Uniq(t.ary), t.expected)
	}
}
