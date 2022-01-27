package subcmd_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestTouch(t *testing.T) {
	assert := assert.New(t)

	touch := &subcmd.TouchCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return nil, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "zoo",
			Category: "foo/bar",
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := touch.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}

func TestTouch_AlreadyExits(t *testing.T) {
	assert := assert.New(t)

	touch := &subcmd.TouchCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			URL: "https://docs.esa.io/posts/1",
		}, nil
	}

	err := touch.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.Equal(errors.New("post already exists: https://docs.esa.io/posts/1"), err)
}
