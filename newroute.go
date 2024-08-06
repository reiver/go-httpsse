package httpsse

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/reiver/go-httpsse/internal/safebufio"
)

type internalRoute struct {
	safeReadWriter safebufio.SafeReadWriter
	waitclosed sync.WaitGroup
	waitserving sync.WaitGroup
}

// NewRoute creates and returns a new Route.
//
// Example usage:
//
//	var route httpsse.Route = httpsse.NewRoute()
func NewRoute() Route {
	var route internalRoute

	route.waitclosed.Add(1)
	route.waitserving.Add(1)

	return &route
}

func (receiver *internalRoute) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.safeReadWriter.Set(nil)
	receiver.waitclosed.Done()
	return nil
}

func (receiver *internalRoute) PublishEvent(fn func(EventWriter)error) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == fn {
		return errNilFunction
	}

	receiver.waitserving.Wait()

	err := receiver.safeReadWriter.Write(func(writer io.Writer)error{
		eventwriter := internalEventWriter{writer}

		err := fn(eventwriter)

		var buffer [1]byte = [1]byte{'\n'}
		writer.Write(buffer[:])

		return err
	})
	return err
}

func (receiver *internalRoute) ServeHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	if nil == receiver {
		return
	}
	if nil == responsewriter {
		return
	}

	if nil == request {
		const code int = http.StatusInternalServerError
		var text string = http.StatusText(code)
		http.Error(responsewriter, text, code)
		return
	}

	var netconn net.Conn
	var bufrw *bufio.ReadWriter
	{
		hijacker, casted := responsewriter.(http.Hijacker)
		if !casted {
			const code int = http.StatusInternalServerError
			var text string = http.StatusText(code)
			http.Error(responsewriter, text, code)
			return
		}

		var err error
		netconn, bufrw, err = hijacker.Hijack()
		if nil != err {
			const code int = http.StatusInternalServerError
			var text string = http.StatusText(code)
			http.Error(responsewriter, text, code)
			return
		}
		if nil == netconn {
			const code int = http.StatusInternalServerError
			var text string = http.StatusText(code)
			http.Error(responsewriter, text, code)
			return
		}
		if nil == bufrw {
			const code int = http.StatusInternalServerError
			var text string = http.StatusText(code)
			http.Error(responsewriter, text, code)
			return
		}
	}
	defer netconn.Close()

	{
		_, err := io.WriteString(bufrw,
			"HTTP/1.1 200 OK"                          +"\r\n"+
			"Content-Type: text/event-stream"          +"\r\n"+
			"Date: " + time.Now().Format(time.RFC1123) +"\r\n"+
			"Connection: close"                        +"\r\n"+
			""                                         +"\r\n",
		)
		if nil != err {
			return
		}
		bufrw.Flush()
	}

	receiver.safeReadWriter.Set(bufrw)
	receiver.waitserving.Done()

	receiver.waitclosed.Wait()
}
