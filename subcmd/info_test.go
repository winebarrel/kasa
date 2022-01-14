package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	info := &subcmd.InfoCmd{
		PostNum: 1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetFromPageNum = func(postNum int) (*model.Post, error) {
		assert.Equal(1, postNum)

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
