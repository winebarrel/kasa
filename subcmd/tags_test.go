package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestTags(t *testing.T) {
	assert := assert.New(t)

	tags := &subcmd.TagsCmd{
		Page: 1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetTags = func(postNum int) (*model.Tags, bool, error) {
		assert.Equal(1, postNum)

		return &model.Tags{
			Tags: []struct {
				Name       string "json:\"name\""
				PostsCount int    "json:\"posts_count\""
			}{
				{"foo", 1},
				{"bar", 2},
			},
		}, false, nil
	}

	err := tags.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`        1  foo
        2  bar
`, printer.String())
}

func TestTags_HasNext(t *testing.T) {
	assert := assert.New(t)

	tags := &subcmd.TagsCmd{
		Page: 1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetTags = func(postNum int) (*model.Tags, bool, error) {
		assert.Equal(1, postNum)

		return &model.Tags{
			Tags: []struct {
				Name       string "json:\"name\""
				PostsCount int    "json:\"posts_count\""
			}{
				{"foo", 1},
				{"bar", 2},
			},
		}, true, nil
	}

	err := tags.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`        1  foo
        2  bar
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}
