package errors

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrUnexpectedEOF = erorr.Error("httpsse: unexpected end-of-file (eof)")
)

const (
	ErrNilRuneScanner = erorr.Error("httpsse: nil rune-scanner")
	ErrNilWriter      = erorr.Error("httpsse: nil writer")
)
