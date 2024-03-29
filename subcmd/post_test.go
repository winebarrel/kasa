package subcmd_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestPost_New(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := os.CreateTemp("", "bodyMd")
	bodyFile.WriteString("bodyMd") //nolint:errcheck
	defer os.Remove(bodyFile.Name())

	post := &subcmd.PostCmd{
		Name:     "zoo",
		Body:     bodyFile.Name(),
		Category: "foo/bar",
		Wip:      false,
		Tags:     []string{"tagA", "tagB"},
		Message:  "msg",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "zoo",
			BodyMd:   "bodyMd",
			Tags:     []string{"tagA", "tagB"},
			Category: "foo/bar",
			Wip:      esa.Bool(false),
			Message:  "msg",
		}, newPostBody)

		assert.Equal(0, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/1", nil
	}

	err := post.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1\n", printer.String())
}

func TestPost_Update(t *testing.T) {
	assert := assert.New(t)

	post := &subcmd.PostCmd{
		Path:     "foo/bar/zoo",
		Category: "foo/bar",
		Wip:      false,
		Tags:     []string{"tagA", "tagB"},
		Message:  "msg",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			Number: 1,
			BodyMd: "body",
		}, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "",
			BodyMd:   "",
			Tags:     []string{"tagA", "tagB"},
			Category: "foo/bar",
			Wip:      esa.Bool(false),
			Message:  "msg",
		}, newPostBody)

		assert.Equal(1, postNum)
		assert.False(notice)

		return "https://docs.esa.io/posts/2", nil
	}

	err := post.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/2\n", printer.String())
}
