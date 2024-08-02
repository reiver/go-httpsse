package endofline

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNotEndOfLine = erorr.Error("httpsse: not end-of-line")
)

const (
	errNilRuneScanner = erorr.Error("httpsse: nil rune-scanner")
	errNilWriter      = erorr.Error("httpsse: nil writer")
)
