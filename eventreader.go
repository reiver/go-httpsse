package httpsse

type EventReader interface {
	EventData() string
	EventID() string
	EventName() string
}
