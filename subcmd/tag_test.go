package subcmd_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestTag_Tag(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "foo/bar/",
		Tags:      []string{"bar", "baz"},
		Override:  false,
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

	driver.MockTag = func(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
		switch postNum {
		case 1:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"bar", "baz", "tagA", "tagB"},
			}, tagPostBody)
		case 2:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"bar", "baz", "tagA", "tagB"},
			}, tagPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`tag '[#bar,#baz,#tagA,#tagB]' 'foo/bar/zoo'
tag '[#bar,#baz,#tagA,#tagB]' 'foo/bar/baz'
`, printer.String())
}

func TestTag_Tag_HasMore(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "foo/bar/",
		Tags:      []string{"bar", "baz"},
		Override:  false,
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

	driver.MockTag = func(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
		switch postNum {
		case 1:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"bar", "baz", "tagA", "tagB"},
			}, tagPostBody)
		case 2:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"bar", "baz", "tagA", "tagB"},
			}, tagPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`tag '[#bar,#baz,#tagA,#tagB]' 'foo/bar/zoo'
tag '[#bar,#baz,#tagA,#tagB]' 'foo/bar/baz'
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}

func TestTag_Tag_Override(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "foo/bar/",
		Tags:      []string{"bar", "baz"},
		Override:  true,
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

	driver.MockTag = func(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
		switch postNum {
		case 1:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"bar", "baz"},
			}, tagPostBody)
		case 2:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"bar", "baz"},
			}, tagPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`tag '[#bar,#baz]' 'foo/bar/zoo'
tag '[#bar,#baz]' 'foo/bar/baz'
`, printer.String())
}

func TestTag_Tag_Delete(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "foo/bar/",
		Tags:      []string{},
		Override:  false,
		Delete:    true,
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

	driver.MockTag = func(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
		switch postNum {
		case 1:
			assert.Equal(&model.TagPostBody{
				Tags: []string{},
			}, tagPostBody)
		case 2:
			assert.Equal(&model.TagPostBody{
				Tags: []string{},
			}, tagPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`tag '' 'foo/bar/zoo'
tag '' 'foo/bar/baz'
`, printer.String())
}

func TestTag_Tag_DeleteWithTags(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "foo/bar/",
		Tags:      []string{"bar", "tagB"},
		Override:  false,
		Delete:    true,
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
				Tags:     []string{"tagA", "tagB", "bar"},
				Category: "foo/bar",
				Message:  "barMsg",
			},
		}, false, nil
	}

	driver.MockTag = func(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
		switch postNum {
		case 1:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"tagA"},
			}, tagPostBody)
		case 2:
			assert.Equal(&model.TagPostBody{
				Tags: []string{"tagA"},
			}, tagPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`tag '[#tagA]' 'foo/bar/zoo'
tag '[#tagA]' 'foo/bar/baz'
`, printer.String())
}

func TestTag_Tag_NoTagsOptionError(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "foo/bar/",
		Tags:      []string{},
		Override:  false,
		Delete:    false,
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.Equal(errors.New("missing flags: --body=TAGS"), err)
}

func TestTag_Tags(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.TagCmd{
		Path:      "in:foo/bar/",
		Tags:      []string{"bar", "baz"},
		Search:    true,
		Override:  false,
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(path string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("in:foo/bar/", path)
		assert.Equal(1, postNum)
		updatedAt, _ := time.Parse("2006/01/02", "2022/01/15")

		return []*model.Post{
			{
				Number:    1,
				Name:      "zoo",
				BodyMd:    "zooBody",
				Wip:       false,
				Tags:      []string{"tagA", "tagB"},
				Category:  "foo/bar",
				Message:   "zooMsg",
				UpdatedAt: updatedAt,
			},
			{
				Number:    2,
				Name:      "baz",
				BodyMd:    "bazBody",
				Wip:       true,
				Tags:      []string{"tagA", "tagB"},
				Category:  "foo/bar",
				Message:   "barMsg",
				UpdatedAt: updatedAt,
			},
		}, false, nil
	}

	driver.MockTag = func(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
		updatedAt, _ := time.Parse("2006/01/02", "2022/01/15")

		switch postNum {
		case 1:
			assert.Equal(&model.TagPostBody{
				Tags:      []string{"bar", "baz", "tagA", "tagB"},
				UpdatedAt: updatedAt,
			}, tagPostBody)
		case 2:
			assert.Equal(&model.TagPostBody{
				Tags:      []string{"bar", "baz", "tagA", "tagB"},
				UpdatedAt: updatedAt,
			}, tagPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := tag.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`tag '[#bar,#baz,#tagA,#tagB]' 'foo/bar/zoo'
tag '[#bar,#baz,#tagA,#tagB]' 'foo/bar/baz'
`, printer.String())
}
