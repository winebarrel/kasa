package subcmd_test

import (
	"testing"

	"github.com/kanmu/kasa"
	"github.com/kanmu/kasa/subcmd"
	"github.com/stretchr/testify/assert"
)

func TestMvcat(t *testing.T) {
	assert := assert.New(t)

	mvcat := &subcmd.MvcatCmd{
		From: "foo/bar",
		To:   "bar/baz",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockMoveCategory = func(from string, to string) error {
		assert.Equal("foo/bar", from)
		assert.Equal("bar/baz", to)

		return nil
	}

	err := mvcat.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Empty(printer.String())
}
