package httpsse

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilEventSetter        = erorr.Error("httpsse: nil event-setter")
	errNilEventWriter        = erorr.Error("httpsse: nil event-writer")
	errNilFunction           = erorr.Error("httpsse: nil function")
	errNilHTTPRequest        = erorr.Error("httpsse: nil http-request")
	errNilHTTPResponse       = erorr.Error("httpsse: nil http-response")
	errNilHTTPResponseBody   = erorr.Error("httpsse: nil http-response-body")
	errNilHTTPResponseHeader = erorr.Error("httpsse: nil http-response-header")
	errNilHTTPResponseWriter = erorr.Error("httpsse: nil http-response-writer")
	errNilReceiver           = erorr.Error("httpsse: nil receiver")
	errNilRuneScanner        = erorr.Error("httpsse: nil rune-scanner")
	errNilWriter             = erorr.Error("httpsse: nil writer")
)
