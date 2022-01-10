package postname_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/postname"
)

func TestPostnameSplit(t *testing.T) {
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
		cat, name := postname.Split(t.input)
		assert.Equal(t.category, cat)
		assert.Equal(t.name, name)
	}
}

func TestPostnameJoin(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := []struct {
		category string
		name     string
		expected string
	}{
		{"bar/zoo", "foo", "bar/zoo/foo"},
		{"", "foo", "foo"},
		{"bar/zoo", "", "bar/zoo/"},
		{"/bar/zoo", "foo", "bar/zoo/foo"},
	}

	for _, t := range tests {
		assert.Equal(postname.Join(t.category, t.name), t.expected)
	}
}
