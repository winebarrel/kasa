package subcmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestEdit(t *testing.T) {
	assert := assert.New(t)
	dir, _ := os.Getwd()
	t.Setenv("EDITOR", filepath.Join(dir, "../tool/edit.sh"))

	edit := &subcmd.EditCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			BodyMd: "body\r\n",
		}, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			BodyMd: "modified\n",
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := edit.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}

func TestEdit_NotModified(t *testing.T) {
	assert := assert.New(t)
	t.Setenv("EDITOR", "cat")

	edit := &subcmd.EditCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			BodyMd: "body\r\n",
		}, nil
	}

	// driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
	// 	assert.Equal(&model.NewPostBody{
	// 		BodyMd: "body",
	// 	}, newPostBody)

	// 	assert.Equal(0, postNum)
	// 	assert.False(notice)

	// 	return "https://docs.esa.io/posts/1", nil
	// }

	err := edit.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Empty(printer.String())
}

func TestEdit_NewFile(t *testing.T) {
	assert := assert.New(t)
	dir, _ := os.Getwd()
	t.Setenv("EDITOR", filepath.Join(dir, "../tool/edit.sh"))

	edit := &subcmd.EditCmd{
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
			BodyMd:   "modified\n",
			Category: "foo/bar",
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := edit.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}
