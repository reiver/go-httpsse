package httpsse

import (
	"io"
)

type eventWriter interface {
	io.Writer
	SetEventName(string)
	SetEventID(string)
}
