package name

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNotName       = erorr.Error("httpsse: not name")
	ErrUnexpectedEOF = erorr.Error("httpsse: unexpected end-of-file (eof)")
)

const (
	errNilRuneScanner = erorr.Error("httpsse: nil rune-scanner")
	errNilWriter      = erorr.Error("httpsse: nil writer")
)
