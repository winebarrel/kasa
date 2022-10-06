package subcmd_test

import (
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestRm_Dir(t *testing.T) {
	assert := assert.New(t)

	rm := &subcmd.RmCmd{
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
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, false, nil
	}

	driver.MockDelete = func(postNum int) error {
		assert.Contains([]int{1, 2}, postNum)

		return nil
	}

	err := rm.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
`, printer.String())
}

func TestRm_HasMore(t *testing.T) {
	assert := assert.New(t)

	rm := &subcmd.RmCmd{
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
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, true, nil
	}

	driver.MockDelete = func(postNum int) error {
		assert.Contains([]int{1, 2}, postNum)

		return nil
	}

	err := rm.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}

func TestRm_Tag(t *testing.T) {
	assert := assert.New(t)

	rm := &subcmd.RmCmd{
		Path:      "#tagA",
		Force:     true,
		Page:      1,
		Recursive: true,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockListOrTagSearch = func(path string, postNum int, recursive bool) ([]*model.Post, bool, error) {
		assert.Equal("#tagA", path)
		assert.Equal(1, postNum)
		assert.True(recursive)

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

	driver.MockDelete = func(postNum int) error {
		assert.Contains([]int{1, 2}, postNum)

		return nil
	}

	err := rm.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
`, printer.String())
}
