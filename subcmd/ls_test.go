package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestLs(t *testing.T) {
	assert := assert.New(t)

	ls := &subcmd.LsCmd{
		Path:      "foo/bar/",
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
				Name:     "zoo",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
			},
			{
				Name:     "baz",
				Wip:      false,
				Tags:     []string{"tagB"},
				Category: "foo/bar",
			},
		}, false, nil
	}

	err := ls.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`0001-01-01 12:00:00  WIP              foo/bar/zoo  [#tagA,#tagB]
0001-01-01 12:00:00  -                foo/bar/baz  [#tagB]
`, printer.String())
}

func TestLs_HasMore(t *testing.T) {
	assert := assert.New(t)

	ls := &subcmd.LsCmd{
		Path:      "foo/bar/",
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
				Name:     "zoo",
				Wip:      true,
				Tags:     []string{"tagA", "tagB"},
				Category: "foo/bar",
			},
			{
				Name:     "baz",
				Wip:      false,
				Tags:     []string{"tagB"},
				Category: "foo/bar",
			},
		}, true, nil
	}

	err := ls.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`0001-01-01 12:00:00  WIP              foo/bar/zoo  [#tagA,#tagB]
0001-01-01 12:00:00  -                foo/bar/baz  [#tagB]
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}
