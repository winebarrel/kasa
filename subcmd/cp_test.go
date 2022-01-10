package subcmd_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestCp_DirToDir(t *testing.T) {
	assert := assert.New(t)

	cp := &subcmd.CpCmd{
		Source:    "foo/bar/",
		Target:    "bar/baz/",
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockListOrTagSearch = func(path string, postNum int, recursive bool) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", path)
		assert.Equal(1, postNum)
		assert.True(recursive)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				BodyMd:   "zooBody",
				Wip:      false,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "zooMsg",
			},
			{
				Number:   2,
				Name:     "baz",
				BodyMd:   "bazBody",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "barMsg",
			},
		}, false, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int) (string, error) {
		switch newPostBody.Name {
		case "zoo":
			assert.Equal(&model.NewPostBody{
				Name:     "zoo",
				BodyMd:   "zooBody",
				Wip:      false,
				Tags:     []string{"tagA", "tagB"},
				Category: "bar/baz/",
				Message:  "zooMsg",
			}, newPostBody)
		case "baz":
			assert.Equal(&model.NewPostBody{
				Name:     "baz",
				BodyMd:   "bazBody",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "bar/baz/",
				Message:  "barMsg",
			}, newPostBody)
		default:
			assert.Failf("invalid post", "post name=%s", newPostBody.Name)
		}

		assert.Equal(0, postNum)

		return "https://docs.esa.io/posts/0", nil
	}

	err := cp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`cp 'foo/bar/zoo' 'bar/baz/zoo'
https://docs.esa.io/posts/0        bar/baz/zoo
cp 'foo/bar/baz' 'bar/baz/baz'
https://docs.esa.io/posts/0        bar/baz/baz
`, printer.String())
}

func TestCp_HasMore(t *testing.T) {
	assert := assert.New(t)

	cp := &subcmd.CpCmd{
		Source:    "foo/bar/",
		Target:    "bar/baz/",
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockListOrTagSearch = func(path string, postNum int, recursive bool) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", path)
		assert.Equal(1, postNum)
		assert.True(recursive)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				BodyMd:   "zooBody",
				Wip:      false,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "zooMsg",
			},
			{
				Number:   2,
				Name:     "baz",
				BodyMd:   "bazBody",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "barMsg",
			},
		}, true, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int) (string, error) {
		switch newPostBody.Name {
		case "zoo":
			assert.Equal(&model.NewPostBody{
				Name:     "zoo",
				BodyMd:   "zooBody",
				Wip:      false,
				Tags:     []string{"tagA", "tagB"},
				Category: "bar/baz/",
				Message:  "zooMsg",
			}, newPostBody)
		case "baz":
			assert.Equal(&model.NewPostBody{
				Name:     "baz",
				BodyMd:   "bazBody",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "bar/baz/",
				Message:  "barMsg",
			}, newPostBody)
		default:
			assert.Failf("invalid post", "post name=%s", newPostBody.Name)
		}

		assert.Equal(0, postNum)

		return "https://docs.esa.io/posts/0", nil
	}

	err := cp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`cp 'foo/bar/zoo' 'bar/baz/zoo'
https://docs.esa.io/posts/0        bar/baz/zoo
cp 'foo/bar/baz' 'bar/baz/baz'
https://docs.esa.io/posts/0        bar/baz/baz
`+"(has more pages. current page is 1, try `-p 2`)\n", printer.String())
}

func TestCp_FileToFile(t *testing.T) {
	assert := assert.New(t)

	cp := &subcmd.CpCmd{
		Source:    "foo/bar/zoo",
		Target:    "bar/baz/baz",
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockListOrTagSearch = func(path string, postNum int, recursive bool) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/zoo", path)
		assert.Equal(1, postNum)
		assert.True(recursive)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				BodyMd:   "zooBody",
				Wip:      false,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "zooMsg",
			},
		}, false, nil
	}

	driver.MockPost = func(newPostBody *model.NewPostBody, postNum int) (string, error) {
		assert.Equal(&model.NewPostBody{
			Name:     "baz",
			BodyMd:   "zooBody",
			Wip:      false,
			Tags:     []string{"tagA", "tagB"},
			Category: "bar/baz",
			Message:  "zooMsg",
		}, newPostBody)

		assert.Equal(0, postNum)

		return "https://docs.esa.io/posts/0", nil
	}

	err := cp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`cp 'foo/bar/zoo' 'bar/baz/baz'
https://docs.esa.io/posts/0        bar/baz/baz
`, printer.String())
}

func TestCp_DirToFile(t *testing.T) {
	assert := assert.New(t)

	cp := &subcmd.CpCmd{
		Source:    "foo/bar/",
		Target:    "bar/baz/baz",
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockListOrTagSearch = func(path string, postNum int, recursive bool) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", path)
		assert.Equal(1, postNum)
		assert.True(recursive)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				BodyMd:   "zooBody",
				Wip:      false,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "zooMsg",
			},
			{
				Number:   2,
				Name:     "baz",
				BodyMd:   "bazBody",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
				Message:  "barMsg",
			},
		}, false, nil
	}

	err := cp.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.Equal(errors.New("target 'bar/baz/baz' is not a category"), err)
}
