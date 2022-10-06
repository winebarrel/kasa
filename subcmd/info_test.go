package subcmd_test

import (
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/esa/model"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	info := &subcmd.InfoCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			Category: "foo/bar",
			Name:     "name",
			FullName: "full_name",
			BodyMd:   "body",
			BodyHTML: "html",
			CreatedBy: &model.PostAuthor{
				Myself:     false,
				Name:       "creator",
				ScreenName: "cre",
				Icon:       "iconA",
			},
			UpdatedBy: &model.PostAuthor{
				Myself:     true,
				Name:       "updater",
				ScreenName: "upd",
				Icon:       "iconB",
			},
		}, nil
	}

	err := info.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`{
  "number": 0,
  "name": "name",
  "full_name": "full_name",
  "wip": false,
  "created_at": "0001-01-01T00:00:00Z",
  "message": "",
  "url": "",
  "updated_at": "0001-01-01T00:00:00Z",
  "tags": null,
  "category": "foo/bar",
  "revision_number": 0,
  "created_by": {
    "myself": false,
    "name": "creator",
    "screen_name": "cre",
    "icon": "iconA"
  },
  "updated_by": {
    "myself": true,
    "name": "updater",
    "screen_name": "upd",
    "icon": "iconB"
  }
}
`, printer.String())
}
