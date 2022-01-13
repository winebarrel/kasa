package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
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
			BodyMd: "body",
		}, nil
	}

	err := open.Run(&kasa.Context{
		Driver: driver,
	})
	assert.NoError(err)
}
