package httpsse

import (
	"io"
	"net/http"
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
//
//		err := sseServer.PublishEvent(fn)
//
//		// ...
//
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

func (receiver internalServer) PublishEvent(fn func(io.Writer)error) error {

	if nil == fn {
		return errNilFunction
	}

	var writer io.Writer = receiver.writer
	if nil == writer {
		return errNilWriter
	}

	err := fn(writer)
	if nil != err {
		return err
	}

	_, err = io.WriteString(writer, "\n")
	if nil != err {
		return err
	}

	return nil
}
