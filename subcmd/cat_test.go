package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestCat(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	cat := &subcmd.CatCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)
		return &model.Post{}, nil
	}

	cat.Run(&kasa.Context{
		Team:   "example",
		Driver: driver,
	})
}
