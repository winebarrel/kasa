package subcmd_test

import (
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert := assert.New(t)

	search := &subcmd.SearchCmd{
		Query: "foo bar zoo",
		Page:  1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(path string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo bar zoo", path)
		assert.Equal(1, postNum)
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

	err := search.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`0001-01-01 12:00:00  WIP              foo/bar/zoo  [#tagA,#tagB]
0001-01-01 12:00:00  -                foo/bar/baz  [#tagB]
`, printer.String())
}

func TestSearch_HasMore(t *testing.T) {
	assert := assert.New(t)

	search := &subcmd.SearchCmd{
		Query: "foo bar zoo",
		Page:  1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(path string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo bar zoo", path)
		assert.Equal(1, postNum)
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

	err := search.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`0001-01-01 12:00:00  WIP              foo/bar/zoo  [#tagA,#tagB]
0001-01-01 12:00:00  -                foo/bar/baz  [#tagB]
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}

func TestSearch_JSON(t *testing.T) {
	assert := assert.New(t)

	search := &subcmd.SearchCmd{
		Query: "foo bar zoo",
		Json:  true,
		Page:  1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(path string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo bar zoo", path)
		assert.Equal(1, postNum)
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

	err := search.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`{"number":0,"name":"zoo","wip":true,"created_at":"0001-01-01T00:00:00Z","message":"","url":"","updated_at":"0001-01-01T00:00:00Z","tags":["tagA","tagB"],"category":"foo/bar","revision_number":0}
{"number":0,"name":"baz","wip":false,"created_at":"0001-01-01T00:00:00Z","message":"","url":"","updated_at":"0001-01-01T00:00:00Z","tags":["tagB"],"category":"foo/bar","revision_number":0}
`, printer.String())
}
