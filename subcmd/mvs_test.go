package subcmd_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestMvs_FilesToDir(t *testing.T) {
	assert := assert.New(t)

	mvs := &subcmd.MvCmd{
		Source: "foo/bar/",
		Target: "bar/baz/",
		Search: true,
		Force:  true,
		Page:   1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(query string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", query)
		assert.Equal(1, postNum)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, false, nil
	}

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Contains([]int{1, 2}, postNum)
		assert.Equal("bar/baz/", movePostBody.Category)
		assert.Empty(movePostBody.Name)

		return nil
	}

	err := mvs.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/'
mv 'foo/bar/baz' 'bar/baz/'
`, printer.String())
}

func TestMvs_HasMore(t *testing.T) {
	assert := assert.New(t)

	mvs := &subcmd.MvCmd{
		Source: "foo/bar/",
		Target: "bar/baz/",
		Search: true,
		Force:  true,
		Page:   1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(query string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", query)
		assert.Equal(1, postNum)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, true, nil
	}

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Contains([]int{1, 2}, postNum)
		assert.Equal("bar/baz/", movePostBody.Category)
		assert.Empty(movePostBody.Name)

		return nil
	}

	err := mvs.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/'
mv 'foo/bar/baz' 'bar/baz/'
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}

func TestMvs_FileToFile(t *testing.T) {
	assert := assert.New(t)

	mvs := &subcmd.MvCmd{
		Source: "foo/bar/zoo",
		Target: "bar/baz/qux",
		Search: true,
		Force:  true,
		Page:   1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(query string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/zoo", query)
		assert.Equal(1, postNum)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				Category: "foo/bar",
			},
		}, false, nil
	}

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Equal(1, postNum)
		assert.Equal("bar/baz", movePostBody.Category)
		assert.Equal("qux", movePostBody.Name)

		return nil
	}

	err := mvs.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/qux'
`, printer.String())
}

func TestMvs_FilesToFile(t *testing.T) {
	assert := assert.New(t)

	mvs := &subcmd.MvCmd{
		Source: "foo/bar/",
		Target: "bar/baz/qux",
		Search: true,
		Force:  true,
		Page:   1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(query string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", query)
		assert.Equal(1, postNum)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, false, nil
	}

	err := mvs.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.Equal(errors.New("target 'bar/baz/qux' is not a category"), err)
}
