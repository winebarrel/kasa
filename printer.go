package kasa

import "fmt"

type Printer interface {
	Printf(string, ...interface{}) (n int, err error)
	Println(...interface{}) (n int, err error)
}

type PrinterImpl struct {
}

func (*PrinterImpl) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func (*PrinterImpl) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
