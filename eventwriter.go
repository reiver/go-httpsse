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
//
// You can give anything that is an EventWriter to [Client].Decode
//
// [Event] is an EventWriter, and can be passed to [Client].Decode
//
// But you can also use other types, including creating your own, that are EventWriter, and pass it to [Client].Decode
type EventWriter interface {
	io.Writer
	SetEventName(string)
	SetEventID(string)
}
