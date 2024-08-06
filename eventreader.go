package httpsse

type EventReader interface {
	EventData() []string
	EventDataCollapsed() string
	EventID() string
	EventName() string
}
