package httpsse

import (
	"io"
)

type EventWriter interface {
	io.Writer
	SetEventName(string)
	SetEventID(string)
}
