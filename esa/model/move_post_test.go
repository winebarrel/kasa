package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/esa/model"
)

func TestMovePostBodyString(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(nil)

	tests := map[*model.MovePostBody]string{
		{Name: "foo", Category: "bar/zoo"}:  "bar/zoo/foo",
		{Name: "foo", Category: ""}:         "foo",
		{Name: "", Category: "bar/zoo"}:     "bar/zoo/",
		{Name: "foo", Category: "/bar/zoo"}: "bar/zoo/foo",
	}

	for data, expected := range tests {
		assert.Equal(data.String(), expected)
	}
}
