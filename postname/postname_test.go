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
		assert.Equal(t.expected, postname.Join(t.category, t.name))
	}
}

func TestPostnameAppendCategoryN(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := []struct {
		src      string
		extra    string
		n        int
		expected string
	}{
		// Minus
		{"foo/bar/zoo", "hoge/fuga/piyo", 0, "foo/bar/zoo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", -1, "foo/bar/zoo/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", -2, "foo/bar/zoo/fuga/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", -3, "foo/bar/zoo/hoge/fuga/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", -4, "foo/bar/zoo/hoge/fuga/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo/", -1, "foo/bar/zoo/piyo"},
		{"foo/bar/zoo/", "hoge/fuga/piyo", -1, "foo/bar/zoo/piyo"},
		{"foo/bar/zoo", "/", -2, "foo/bar/zoo"},
		{"foo/bar/zoo", "", -2, "foo/bar/zoo"},
		// Plus
		{"foo/bar/zoo", "hoge/fuga/piyo", 1, "foo/bar/zoo/hoge/fuga/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", 2, "foo/bar/zoo/fuga/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", 3, "foo/bar/zoo/piyo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", 4, "foo/bar/zoo"},
		{"foo/bar/zoo", "hoge/fuga/piyo", 5, "foo/bar/zoo"},
		{"foo/bar/zoo", "hoge/fuga/piyo/", 1, "foo/bar/zoo/hoge/fuga/piyo"},
		{"foo/bar/zoo/", "hoge/fuga/piyo", 1, "foo/bar/zoo/hoge/fuga/piyo"},
		{"foo/bar/zoo", "/", 2, "foo/bar/zoo"},
		{"foo/bar/zoo", "", 2, "foo/bar/zoo"},
	}

	for _, t := range tests {
		assert.Equal(t.expected, postname.AppendCategoryN(t.src, t.extra, t.n))
	}
}
