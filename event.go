package httpsse

import (
	"io"
	"strings"
)

// An event represents an HTTP server-sent event.
type Event struct {
	name string
	id string
	data internalField
}

var _ EventReader = &Event{}
var _ EventSetter = &Event{}

func (receiver *Event) AppendEventDatum(value string) {
	receiver.data.Append(value)
}

// Equal returns whether two events are equal.
//
// Two events arer equal if their event-data, event-id, and event-name are all equal.
func (receiver Event) Equal(that EventReader) bool {
	if nil == that {
		return false
	}

	return receiver.EventName()          == that.EventName() &&
	       receiver.EventID()            == that.EventID()   &&
	       receiver.EventDataCollapsed() == that.EventDataCollapsed()
}

func (receiver Event) EventData() []string {
	return receiver.data.Values()
}

// EventID returns the event-ID
//
// For example, if this is the event:
//
//	event: banana
//	id: yellow-123
//	data: once
//	data: twice
//	data: thirce
//	data: fource
//
// Then EventID would return:
//
//	"once\ntwice\nthrice\nfource"
func (receiver Event) EventDataCollapsed() string {
	var data string = receiver.data.String()

	length := len(data)

	if length <= 0 {
		return ""
	}

	return data
}

// EventID returns the event-ID.
//
// For example, if this is the event:
//
//	event: banana
//	id: yellow-123
//	data: once
//	data: twice
//	data: thirce
//	data: fource
//
// Then EventID would return:
//
//	"yellow-123"
func (receiver Event) EventID() string {
	return receiver.id
}

// EventName returns the event-name.
//
// For example, if this is the event:
//
//	event: banana
//	id: yellow-123
//	data: once
//	data: twice
//	data: thirce
//	data: fource
//
// Then EventName would return:
//
//	"banana"
func (receiver Event) EventName() string {
	return receiver.name
}

// Reset resets the event to be empty again.
func (receiver *Event) Reset() {
	receiver.name = ""
	receiver.id = ""
	receiver.data.Reset()
}

// SetEventID sets the event-ID.
func (receiver *Event) SetEventID(value string) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.id = value
}

// SetEventName sets the event-name.
func (receiver *Event) SetEventName(value string) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.name = value
}

// String returns the serialize form of the event, as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
// So, for example, if event-ID is:
//
//	"yellow-123"
//
// And, for example. event-name is"
//
//	"banana"
//
// And the event-data is:
//
//	"once\ntwice\nthrice\nfource"
//
// Then this would return:
//
//	": {\n"+
//	"event: banana\n" +
//	"id: yello123\n"  +
//	"data: once\n"    +
//	"data: twice\n"   +
//	"data: thice\n"   +
//	"data: fource\n"  +
//	": }\n"+
//	"\n"
func (receiver Event) String() string {
	var builder strings.Builder
	receiver.writeto(&builder)

	return builder.String()
}

// WriteTo writers the serialized form of the event, as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
// So, for example, if event-ID is:
//
//	"yellow-123"
//
// And, for example. event-name is"
//
//	"banana"
//
// And the event-data is:
//
//	"once\ntwice\nthrice\nfource"
//
// Then this would write:
//
//	": {\n"+
//	"event: banana\n" +
//	"id: yello123\n"  +
//	"data: once\n"    +
//	"data: twice\n"   +
//	"data: thice\n"   +
//	"data: fource\n"  +
//	": }\n"+
//	"\n"
func (receiver Event) WriteTo(writer io.Writer) (n int64, err error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var builder strings.Builder
	receiver.writeto(&builder)

	{
		written, err := io.WriteString(writer, builder.String())
		n = int64(written)
		return n, err
	}
}

func (receiver Event) writeto(stringwriter io.StringWriter) {
	if nil == stringwriter {
		return
	}

	{
		stringwriter.WriteString(": {"+"\n")
	}

	{
		var eventname string = receiver.EventName()

		if "" != eventname {
			stringwriter.WriteString("event: ")
			stringwriter.WriteString(eventname)
			stringwriter.WriteString("\n")
		}
	}

	{
		var eventid string = receiver.EventID()

		if "" != eventid {
			stringwriter.WriteString("id: ")
			stringwriter.WriteString(eventid)
			stringwriter.WriteString("\n")
		}
	}

	{
		var eventdata string = receiver.EventDataCollapsed()

		if "" != eventdata {
			stringwriter.WriteString("data: ")
			stringwriter.WriteString( strings.ReplaceAll(eventdata, "\n", "\ndata: ") )
			stringwriter.WriteString("\n")
		}
	}

	{
		stringwriter.WriteString(": }"+"\n")
	}

	{
		stringwriter.WriteString("\n")
	}
}
