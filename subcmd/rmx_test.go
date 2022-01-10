package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestRmx_Dir(t *testing.T) {
	assert := assert.New(t)

	rmx := &subcmd.RmxCmd{
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

	err := rmx.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
`, printer.String())
}

func TestRmx_HasMore(t *testing.T) {
	assert := assert.New(t)

	rmx := &subcmd.RmxCmd{
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

	err := rmx.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
`+"(has more pages. current page is 1, try `-p 2`)\n", printer.String())
}

func TestRmx_Tag(t *testing.T) {
	assert := assert.New(t)

	rmx := &subcmd.RmxCmd{
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

	err := rmx.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
`, printer.String())
}
