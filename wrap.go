package httpsse

import (
	"io"
	"net/http"

	"github.com/reiver/go-erorr"
)

// SererWrap is used, usually from ServeHTTP, to turn an http.ResponseWriter to an HTTP server-send event (SSE) server.
//
// For example:
//
//	func ServeHTTP(responsewriter http.ResponseWriter, request *http.Request) {
//		// ...
//
//		sseServer, err := httpsse.ServerWrap(responsewriter)
//
//		// ...
//	}
func ServerWrap(responsewriter http.ResponseWriter) (Server, error) {
	if nil == responsewriter {
		return nil, errNilHTTPResponseWriter
	}

	{
		var header http.Header = responsewriter.Header()
		if nil == header {
			return nil, errNilHTTPResponseHeader
		}

		header.Add("Content-Type", ContentType)
	}

	responsewriter.WriteHeader(http.StatusOK)

	return serverwrap(responsewriter)
}

func serverwrap(writer io.Writer) (Server, error) {
	if nil == writer {
		return nil, errNilWriter
	}

	var server = internalServer{
		writer:writer,
	}

	return server, nil
}

type internalServer struct {
	writer io.Writer
}

func (receiver internalServer) PublishEvent(src any) error {

	var writer io.Writer = receiver.writer
	if nil == writer {
		return errNilWriter
	}

	switch event := src.(type) {
	case io.WriterTo:
		_, err := event.WriteTo(writer)
		if nil != err {
			return erorr.Errorf("httpsse: cannot write-to writer: %w", err)
		}
	case EventReader:
		var ev Event
		ev.SetEventID(event.EventID())
		ev.SetEventName(event.EventName())
		io.WriteString(&ev, event.EventData())

		_, err := ev.WriteTo(writer)
		if nil != err {
			return erorr.Errorf("httpsse: cannot write event to writer: %w", err)
		}
	default:
		return erorr.Errorf("httpsse: cannot publish something of type %T", src)
	}

	return nil
}
