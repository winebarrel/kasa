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

func TestComment(t *testing.T) {
	assert := assert.New(t)
	bodyFile, _ := ioutil.TempFile("", "bodyMd")
	bodyFile.WriteString("bodyMd")
	defer os.Remove(bodyFile.Name())

	append := &subcmd.CommentCmd{
		Path: "foo/bar/zoo",
		Body: bodyFile.Name(),
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			Number: 1,
		}, nil
	}

	driver.MockComment = func(newPostBody *model.NewCommentBody, postNum int) (string, error) {
		assert.Equal(&model.NewCommentBody{
			BodyMd: "bodyMd",
		}, newPostBody)

		assert.Equal(1, postNum)

		return "https://docs.esa.io/posts/1#comment-1234567", nil
	}

	err := append.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Equal("https://docs.esa.io/posts/1#comment-1234567\n", printer.String())
}
