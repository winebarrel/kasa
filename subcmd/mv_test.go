package subcmd_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestMv_DirToDir(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
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

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/'
mv 'foo/bar/baz' 'bar/baz/'
`, printer.String())
}

func TestMv_HasMore(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
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

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/'
mv 'foo/bar/baz' 'bar/baz/'
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}

func TestMv_TagToDir(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
		Source:    "#tagA",
		Target:    "bar/baz/",
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

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Contains([]int{1, 2}, postNum)
		assert.Equal("bar/baz/", movePostBody.Category)
		assert.Empty(movePostBody.Name)

		return nil
	}

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/'
mv 'foo/bar/baz' 'bar/baz/'
`, printer.String())
}

func TestMv_FileToFile(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
		Source:    "foo/bar/zoo",
		Target:    "bar/baz/zoo",
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
				Category: "foo/bar",
			},
		}, false, nil
	}

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Equal(1, postNum)
		assert.Equal("bar/baz", movePostBody.Category)
		assert.Equal("zoo", movePostBody.Name)

		return nil
	}

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/zoo'
`, printer.String())
}

func TestMv_DirToFile(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
		Source:    "foo/bar/",
		Target:    "bar/baz/zoo",
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

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.Equal(errors.New("target 'bar/baz/zoo' is not a category"), err)
}

func TestMv_WithCat_Minus(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
		Source:    "foo/bar/",
		Target:    "bar/baz/",
		Force:     true,
		WithCat:   -1,
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
				Category: "foo/bar/hoge",
			},
		}, false, nil
	}

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Contains([]int{1, 2}, postNum)
		assert.Contains([]string{"bar/baz/bar", "bar/baz/hoge"}, movePostBody.Category)
		assert.Empty(movePostBody.Name)

		return nil
	}

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/bar/'
mv 'foo/bar/hoge/baz' 'bar/baz/hoge/'
`, printer.String())
}

func TestMv_WithCat_Plus(t *testing.T) {
	assert := assert.New(t)

	mv := &subcmd.MvCmd{
		Source:    "foo/bar/",
		Target:    "bar/baz/",
		Force:     true,
		WithCat:   1,
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
				Category: "foo/bar/hoge",
			},
		}, false, nil
	}

	driver.MockMove = func(movePostBody *model.MovePostBody, postNum int) error {
		assert.Contains([]int{1, 2}, postNum)
		assert.Contains([]string{"bar/baz/foo/bar", "bar/baz/foo/bar/hoge"}, movePostBody.Category)
		assert.Empty(movePostBody.Name)

		return nil
	}

	err := mv.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`mv 'foo/bar/zoo' 'bar/baz/foo/bar/'
mv 'foo/bar/hoge/baz' 'bar/baz/foo/bar/hoge/'
`, printer.String())
}
