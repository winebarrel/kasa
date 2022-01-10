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
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)
		return &model.Post{BodyMd: "body"}, nil
	}

	err := cat.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("body\n", printer.String())
}
