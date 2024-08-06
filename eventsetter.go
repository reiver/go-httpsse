package httpsse

// EventSetter represents an event that can be written to.
//
// Use Write to write the event-data.
//
// Use SetEventName to set the event-name.
//
// Use SetEventID to set the event-ID.
//
// You can give anything that is an EventSetter to [Client].Decode
//
// [Event] is an EventSetter, and can be passed to [Client].Decode
//
// But you can also use other types, including creating your own, that are EventSetter, and pass it to [Client].Decode
type EventSetter interface {
	AppendEventDatum(string)
	SetEventName(string)
	SetEventID(string)
}
