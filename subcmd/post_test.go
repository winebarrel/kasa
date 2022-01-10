package subcmd_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestPost_New(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := ioutil.TempFile("", "bodyMd")
	bodyFile.WriteString("bodyMd")
	defer os.Remove(bodyFile.Name())

	post := &subcmd.PostCmd{
		Name:     "zoo",
		Body:     bodyFile.Name(),
		Category: "foo/bar",
		WIP:      false,
		Tags:     []string{"tagA", "tagB"},
		Message:  "msg",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "zoo",
			BodyMd:   "bodyMd",
			Tags:     []string{"tagA", "tagB"},
			Category: "foo/bar",
			WIP:      false,
			Message:  "msg",
		}, newPostBody)

		assert.Equal(0, postNum)

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
		PostNum:  1,
		Category: "foo/bar",
		WIP:      false,
		Tags:     []string{"tagA", "tagB"},
		Message:  "msg",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "",
			BodyMd:   "",
			Tags:     []string{"tagA", "tagB"},
			Category: "foo/bar",
			WIP:      false,
			Message:  "msg",
		}, newPostBody)

		assert.Equal(1, postNum)

		return "https://docs.esa.io/posts/2", nil
	}

	err := post.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/2\n", printer.String())
}
