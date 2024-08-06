package safebufio

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilBufferedReadWriter = erorr.Error("httpsse: nil buffered read-writer")
	errNilReceiver           = erorr.Error("httpsse: nil receiver")
)
