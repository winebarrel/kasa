package subcmd_test

import (
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
	"github.com/winebarrel/kasa/utils"
)

func TestOpen(t *testing.T) {
	assert := assert.New(t)

	open := &subcmd.OpenCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			URL: "https://docs.esa.io/posts/1",
		}, nil
	}

	monkey.Patch(utils.OpenInBrowser, func(u string) error {
		assert.Equal("https://docs.esa.io/posts/1", u)
		return nil
	})

	err := open.Run(&kasa.Context{
		Driver: driver,
	})

	assert.NoError(err)
}

func TestOpen_Category(t *testing.T) {
	assert := assert.New(t)

	open := &subcmd.OpenCmd{
		Path: "foo/bar/",
	}

	monkey.Patch(utils.OpenInBrowser, func(u string) error {
		assert.Equal("https://docs.esa.io/#path=foo%2Fbar%2F", u)
		return nil
	})

	err := open.Run(&kasa.Context{
		Team: "docs",
	})

	assert.NoError(err)
}
