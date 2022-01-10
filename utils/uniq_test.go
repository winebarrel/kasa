package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/utils"
)

func TestLs(t *testing.T) {
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
