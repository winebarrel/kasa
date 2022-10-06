package kasa

import (
	"github.com/kanmu/kasa/esa"
)

type Context struct {
	Team   string
	Driver esa.Driver
	Fmt    Printer
}
