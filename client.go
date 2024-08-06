package httpsse

// Client represents the client side of an HTTP SSE connection.
type Client interface {
	Close() error
	Decode(interface{}) error
	Err() error
	Next() bool
}
