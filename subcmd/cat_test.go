package subcmd_test

import (
	"errors"
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestCat(t *testing.T) {
	assert := assert.New(t)

	cat := &subcmd.CatCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			BodyMd: "body",
		}, nil
	}

	err := cat.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("body", printer.String())
}

func TestCat_NotFound(t *testing.T) {
	assert := assert.New(t)

	cat := &subcmd.CatCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return nil, nil
	}

	err := cat.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.Equal(errors.New("post not found"), err)
}

func TestCat_URL(t *testing.T) {
	assert := assert.New(t)

	cat := &subcmd.CatCmd{
		Path: "https://docs.esa.io/posts/1",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetFromPageNum = func(postNum int) (*model.Post, error) {
		assert.Equal(1, postNum)

		return &model.Post{
			BodyMd: "body",
		}, nil
	}

	err := cat.Run(&kasa.Context{
		Team:   "docs",
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("body", printer.String())
}

func TestCat_NUM(t *testing.T) {
	assert := assert.New(t)

	cat := &subcmd.CatCmd{
		Path: "//2",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetFromPageNum = func(postNum int) (*model.Post, error) {
		assert.Equal(2, postNum)

		return &model.Post{
			BodyMd: "body",
		}, nil
	}

	err := cat.Run(&kasa.Context{
		Team:   "docs",
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("body", printer.String())
}
