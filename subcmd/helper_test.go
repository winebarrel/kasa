package subcmd_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/esa/model"
)

type MockDriverImpl struct {
	MockGet             func(string) (*model.Post, error)
	MockGetFromPageNum  func(int) (*model.Post, error)
	MockList            func(string, int, bool) ([]*model.Post, bool, error)
	MockSearch          func(string, int) ([]*model.Post, bool, error)
	MockListOrTagSearch func(string, int, bool) ([]*model.Post, bool, error)
	MockPost            func(*model.NewPostBody, int, bool) (string, error)
	MockMove            func(*model.MovePostBody, int, bool) error
	MockMoveCategory    func(string, string) error
	MockDelete          func(int) error
	MockTag             func(*model.TagPostBody, int, bool) error
	MockComment         func(*model.NewCommentBody, int) (string, error)
	MockGetTags         func(int) (*model.Tags, bool, error)
}

func (dri *MockDriverImpl) Get(path string) (*model.Post, error) {
	return dri.MockGet(path)
}

func (dri *MockDriverImpl) GetFromPageNum(pageNum int) (*model.Post, error) {
	return dri.MockGetFromPageNum(pageNum)
}

func (dri *MockDriverImpl) List(path string, pageNum int, recursive bool) ([]*model.Post, bool, error) {
	return dri.MockList(path, pageNum, recursive)
}

func (dri *MockDriverImpl) Search(queryString string, pageNum int) ([]*model.Post, bool, error) {
	return dri.MockSearch(queryString, pageNum)
}

func (dri *MockDriverImpl) ListOrTagSearch(path string, pageNum int, recursive bool) ([]*model.Post, bool, error) {
	return dri.MockListOrTagSearch(path, pageNum, recursive)
}

func (dri *MockDriverImpl) Post(newPostBody *model.NewPostBody, postNum int, notice bool) (string, error) {
	return dri.MockPost(newPostBody, postNum, notice)
}

func (dri *MockDriverImpl) Move(movePostBody *model.MovePostBody, postNum int, notice bool) error {
	return dri.MockMove(movePostBody, postNum, notice)
}

func (dri *MockDriverImpl) MoveCategory(from string, to string) error {
	return dri.MockMoveCategory(from, to)
}

func (dri *MockDriverImpl) Delete(postNum int) error {
	return dri.MockDelete(postNum)
}

func (dri *MockDriverImpl) Tag(tagPostBody *model.TagPostBody, postNum int, notice bool) error {
	return dri.MockTag(tagPostBody, postNum, notice)
}

func (dri *MockDriverImpl) Comment(newCommentBody *model.NewCommentBody, postNum int) (string, error) {
	return dri.MockComment(newCommentBody, postNum)
}

func (dri *MockDriverImpl) GetTags(postNum int) (*model.Tags, bool, error) {
	return dri.MockGetTags(postNum)
}

func NewMockDriver(t *testing.T) *MockDriverImpl {
	t.Helper()
	assert := assert.New(t)

	return &MockDriverImpl{
		MockGet: func(string) (*model.Post, error) {
			assert.Fail("not implemented")
			return nil, nil
		},
		MockGetFromPageNum: func(int) (*model.Post, error) {
			assert.Fail("not implemented")
			return nil, nil
		},
		MockList: func(string, int, bool) ([]*model.Post, bool, error) {
			assert.Fail("not implemented")
			return nil, false, nil
		},
		MockSearch: func(string, int) ([]*model.Post, bool, error) {
			assert.Fail("not implemented")
			return nil, false, nil
		},
		MockListOrTagSearch: func(string, int, bool) ([]*model.Post, bool, error) {
			assert.Fail("not implemented")
			return nil, false, nil
		},
		MockPost: func(*model.NewPostBody, int, bool) (string, error) {
			assert.Fail("not implemented")
			return "", nil
		},
		MockMove: func(*model.MovePostBody, int, bool) error {
			assert.Fail("not implemented")
			return nil
		},
		MockMoveCategory: func(string, string) error {
			assert.Fail("not implemented")
			return nil
		},
		MockDelete: func(int) error {
			assert.Fail("not implemented")
			return nil
		},
		MockTag: func(*model.TagPostBody, int, bool) error {
			assert.Fail("not implemented")
			return nil
		},
		MockComment: func(*model.NewCommentBody, int) (string, error) {
			assert.Fail("not implemented")
			return "", nil
		},
		MockGetTags: func(int) (*model.Tags, bool, error) {
			assert.Fail("not implemented")
			return nil, false, nil
		},
	}
}

type MockPrinterImpl struct {
	buf strings.Builder
}

func (printer *MockPrinterImpl) Printf(format string, a ...interface{}) (n int, err error) {
	printer.buf.WriteString(fmt.Sprintf(format, a...))
	return 0, nil
}

func (printer *MockPrinterImpl) Println(a ...interface{}) (n int, err error) {
	printer.buf.WriteString(fmt.Sprintln(a...))
	return n, nil
}

func (printer *MockPrinterImpl) String() string {
	return printer.buf.String()
}
