package httpsse

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilEventSetter        = erorr.Error("httpsse: nil event-setter")
	errNilFunction           = erorr.Error("httpsse: nil function")
	errNilHTTPResponseHeader = erorr.Error("httpsse: nil http-response-header")
	errNilHTTPResponseWriter = erorr.Error("httpsse: nil http-response-writer")
	errNilReceiver           = erorr.Error("httpsse: nil receiver")
	errNilRuneScanner        = erorr.Error("httpsse: nil rune-scanner")
	errNilWriter             = erorr.Error("httpsse: nil writer")
)
