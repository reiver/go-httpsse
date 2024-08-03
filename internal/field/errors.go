package field

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNotFieldBecauseNotName = erorr.Error("httpsse: not field because no name")
)

const (
	errNilNameWriter   = erorr.Error("httpsse: nil name writer")
	errNilValueWriter  = erorr.Error("httpsse: nil value writer")
)
