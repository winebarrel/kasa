package kasa

import (
	"github.com/winebarrel/kasa/esa"
)

type Context struct {
	Driver esa.Driver
	Fmt    Printer
}
