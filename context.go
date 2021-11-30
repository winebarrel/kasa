package kasa

import (
	"github.com/winebarrel/kasa/esa"
)

type Context struct {
	Team   string
	Driver *esa.Driver
}
