package subcmd_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestAppend(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := os.CreateTemp("", "bodyMd")
	bodyFile.WriteString("bodyMd") //nolint:errcheck
	defer os.Remove(bodyFile.Name())

	append := &subcmd.AppendCmd{
		Path: "foo/bar/zoo",
		Body: bodyFile.Name(),
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			BodyMd: "body",
		}, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			BodyMd: "body\nbodyMd",
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := append.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}

func TestAppend_WithPrefix(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := os.CreateTemp("", "bodyMd")
	bodyFile.WriteString("bodyMd") //nolint:errcheck
	defer os.Remove(bodyFile.Name())

	append := &subcmd.AppendCmd{
		Path:   "foo/bar/zoo",
		Body:   bodyFile.Name(),
		Prefix: "# prefix",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			BodyMd: "body",
		}, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			BodyMd: "body\n# prefix\nbodyMd",
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := append.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}
