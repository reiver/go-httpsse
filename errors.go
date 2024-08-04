package httpsse

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilEventWriter        = erorr.Error("httpsse: nil event-writer")
	errNilHTTPResponseHeader = erorr.Error("httpsse: nil http-response-header")
	errNilHTTPResponseWriter = erorr.Error("httpsse: nil http-response-writer")
	errNilReceiver           = erorr.Error("httpsse: nil receiver")
	errNilRuneScanner        = erorr.Error("httpsse: nil rune-scanner")
	errNilWriter             = erorr.Error("httpsse: nil writer")
)
