package httpsse

import (
	"io"
	"net/http"
)

// Route represents a route to send HTTP server-send events (SSEs).
type Route interface {
	io.Closer
	http.Handler
	PublishEvent(func(io.Writer)error) error
}
