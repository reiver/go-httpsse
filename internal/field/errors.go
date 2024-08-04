package field

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNotField = erorr.Error("httpsse: not field")
)

const (
	errNilNameWriter   = erorr.Error("httpsse: nil name writer")
	errNilValueWriter  = erorr.Error("httpsse: nil value writer")
)
