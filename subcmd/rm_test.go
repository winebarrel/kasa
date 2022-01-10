package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/subcmd"
)

func TestRm(t *testing.T) {
	assert := assert.New(t)

	rm := &subcmd.RmCmd{
		PostNum: 1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockDelete = func(postNum int) error {
		assert.Equal(1, postNum)

		return nil
	}

	err := rm.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)
	assert.Empty(printer.String())
}
