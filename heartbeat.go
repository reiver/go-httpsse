package httpsse

import (
	"fmt"
	"io"
	"time"

	"github.com/reiver/go-lck"
)

// HeatBeat sends a heartbeat comment to the 'route' evey 'wait' duration of time.
//
// You want to use this to make sure the HTTP connection does NOT get closed due to inactivity.
//
// You should almost always use this.
//
// Example usage:
//
//	func ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//
//		// ...
//
//		var route httpsse.Route = httpsse.NewRoute()
//
//		httpsse.HeatBeat(4 * time.Second, route)
//
//		// ...
//
//		route.ServeHTTP(resp, req)
//	}
func HeartBeat(wait time.Duration, route Route) io.Closer {
	if nil == route {
		return nil
	}

	var heartbeat = internalHeartBeat{
		wait:wait,
	}

	heartbeat.spawn(route)

	return &heartbeat
}


type internalHeartBeat struct {
	closed lck.Lockable[bool]
	wait time.Duration
}

func (receiver *internalHeartBeat) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.closed.Set(true)
	return  nil
}

func (receiver *internalHeartBeat) spawn(route Route) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	go receiver.loop(route)
}

func (receiver *internalHeartBeat) loop(route Route) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	var wait time.Duration = receiver.wait

	for !receiver.closed.Get() {
		thump(route, wait)

		time.Sleep(wait)
	}
}

func thump(route Route, wait time.Duration) {

	if nil == route {
		return
	}

	err := route.PublishEvent(func(eventwriter EventWriter)error{
		if nil == eventwriter {
			return errNilEventWriter
		}

		var now time.Time = time.Now()

		{
			var str string = fmt.Sprintf(" heartbeat thump! at unix-time %d (%s)", now.Unix(), now.Format(time.RFC3339))
			eventwriter.WriteComment(str)
		}

		{
			var next time.Time = now.Add(wait)

			var str string = fmt.Sprintf("   next should be at unix-time %d (%s)", next.Unix(), next.Format(time.RFC3339))
			eventwriter.WriteComment(str)
		}

                return nil
	})

	_ = err
}
