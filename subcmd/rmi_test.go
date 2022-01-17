package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestRmi(t *testing.T) {
	assert := assert.New(t)

	rmi := &subcmd.RmiCmd{
		Path:  "foo/bar/zoo",
		Force: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			Number:   1,
			Name:     "zoo",
			Category: "foo/bar",
			BodyMd:   "body",
		}, nil
	}

	driver.MockDelete = func(postNum int) error {
		assert.Equal(1, postNum)

		return nil
	}

	err := rmi.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("rm 'foo/bar/zoo'\n", printer.String())
}
