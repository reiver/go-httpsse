package httpsse

import (
	"io"
)

// EventWriter represents an event that can be written to.
//
// Use Write to write the event-data.
//
// Use SetEventName to set the event-name.
//
// Use SetEventID to set the event-ID.
type EventWriter interface {
	io.Writer
	SetEventName(string)
	SetEventID(string)
}
