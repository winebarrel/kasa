package subcmd_test

import (
	"os"
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := os.CreateTemp("", "bodyMd")
	bodyFile.WriteString("bodyMd") //nolint:errcheck
	defer os.Remove(bodyFile.Name())

	imp := &subcmd.ImportCmd{
		File: bodyFile.Name(),
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "zoo",
			Category: "foo/bar",
			BodyMd:   "bodyMd",
			Wip:      esa.Bool(false),
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := imp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}

func TestImport_WithoutName(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := os.CreateTemp("", "bodyMd")
	bodyFile.WriteString("bodyMd") //nolint:errcheck
	defer os.Remove(bodyFile.Name())

	imp := &subcmd.ImportCmd{
		File: bodyFile.Name(),
		Path: "foo/bar/",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Regexp(`^bodyMd`, newPostBody.Name)
		assert.Equal("foo/bar/", newPostBody.Category)
		assert.Equal("bodyMd", newPostBody.BodyMd)
		assert.False(*newPostBody.Wip)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := imp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}

func TestImport_WIP(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := os.CreateTemp("", "bodyMd")
	bodyFile.WriteString("bodyMd") //nolint:errcheck
	defer os.Remove(bodyFile.Name())

	imp := &subcmd.ImportCmd{
		File: bodyFile.Name(),
		Path: "foo/bar/zoo",
		Wip:  true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "zoo",
			Category: "foo/bar",
			BodyMd:   "bodyMd",
			Wip:      esa.Bool(true),
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := imp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}
