# go-httpsse

Package **httpsse** implements HTTP **server-sent events**, for the Go programming language.

As defined by:
https://html.spec.whatwg.org/multipage/server-sent-events.html

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-httpsse

[![GoDoc](https://godoc.org/github.com/reiver/go-httpsse?status.svg)](https://godoc.org/github.com/reiver/go-httpsse)

## Examples

Here is an example HTTP SSE _server_:

```golang
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/reiver/go-httpsse"
)

func main() {
	var handler http.Handler = http.HandlerFunc(ServeHTTP)

	err := http.ListenAndServe(":8080", handler)
	if nil != err {
		panic(err)
	}
}

func ServeHTTP(responsewriter http.ResponseWriter, request *http.Request) {

	var route httpsse.Route = httpsse.NewRoute()

	// Send a heartbeat comment every 4.567 seconds.
	httpsse.HeartBeat(4567 * time.Millisecond, route)

	// Spawn this into its own go-routine.
	//
	// Having the things writing to the route run in a different go-routine is important
	// so that the call doesn't block before route.ServeHTTP() is called.
	go loop(route)

	route.ServeHTTP(responsewriter, request)
}

// This function isn't important for this example.
//
// Your own functions would do something useful.
func loop(route httpsse.Route) {
	for {
		err := route.PublishEvent(func(eventwriter httpsse.EventWriter)error{

			if nil == eventwriter {
				panic("nil event-writer")
			}

			var eventName string = "status"
			var eventID string = fmt.Sprintf("status-%d", time.Now().Unix())
			var eventData string = "Hello world!"

			eventwriter.WriteEvent(eventName)
			eventwriter.WriteID(eventID)
			eventwriter.WriteData(eventData)

			return nil
		})
		if nil != err {
			fmt.Printf("PUBLISH-EVENT-ERROR: %s \n", err)
		}

		time.Sleep(2351 * time.Millisecond)
	}
}
```

Here is an example HTTP SSE _client_:

```golang
package main

import (
	"fmt"

	"github.com/reiver/go-httpsse"
)

func main() {
	const url string = "http://example.com/api/streaming" // REPLACE THIS WITH THE ACTUAL URL.

	client, err := httpsse.Dial(url)
	if nil != err {
		fmt.Printf("ERROR: had problem dialing %q: %s \n", url, err)
		return
	}
	if nil == client {
		fmt.Println("ERROR: nil client")
		return
	}

	for client.Next() {
		var event httpsse.Event
		err := client.Decode(&event)
		if nil != err {
			fmt.Printf("ERROR: had problem trying to decode the event: %s", err)
			continue
		}

		// You would probably do something useful once you had the event.
		fmt.Println("EVENT:\n", event)
	}
	if err := client.Err(); nil != err {
		fmt.Printf("CLIENT-ERROR: %s", err)
		return
	}
}

```

## Import

To import package **httpsse** use `import` code like the follownig:
```
import "github.com/reiver/go-httpsse"
```

## Installation

To install package **httpsse** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-httpsse
```

## Author

Package **httpsse** was written by [Charles Iliya Krempeaux](http://reiver.link)
