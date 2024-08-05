package httpsse

import (
	"io"
)

// Server represents a HTTP server-send event (SSE) server.
//
// This is more-or-less equivalent to a HTTTP connection that is speaking SSE.
type Server interface {
	PublishEvent(func(io.Writer)error) error
}
