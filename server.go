package httpsse

// Server represents a HTTP server-send event (SSE) server.
//
// This is more-or-less equivalent to a HTTTP connection that is speaking SSE.
type Server interface {
	PublishEvent(any) error
}
