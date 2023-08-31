package subcmd_test

import (
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestUnwip(t *testing.T) {
	assert := assert.New(t)

	unwip := &subcmd.UnwipCmd{
		Path:      "foo/bar/",
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

	driver.MockWip = func(wipPostBody *model.WipPostBody, postNum int, notice bool) error {
		switch postNum {
		case 2:
			assert.Equal(&model.WipPostBody{
				Wip: false,
			}, wipPostBody)
		default:
			assert.Failf("invalid post", "post num=%d", postNum)
		}

		assert.False(notice)

		return nil
	}

	err := unwip.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal("unwip 'foo/bar/baz'\n", printer.String())
}

func TestUnwip_HasMore(t *testing.T) {
	assert := assert.New(t)

	tag := &subcmd.UnwipCmd{
		Path:      "foo/bar/",
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

	driver.MockWip = func(wipPostBody *model.WipPostBody, postNum int, notice bool) error {
		switch postNum {
		case 2:
			assert.Equal(&model.WipPostBody{
				Wip: false,
			}, wipPostBody)
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

	assert.Equal(`unwip 'foo/bar/baz'
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}
