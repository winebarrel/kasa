package subcmd_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestTags(t *testing.T) {
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
