package httpsse

import (
	"io"
	"net/http"

	"github.com/reiver/go-errhttp"
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-utf8"
)

func DialURL(url string) (Client, error) {
	var resp *http.Response
	{
		var err error

		resp, err = http.Get(url)
		if nil != err {
			return nil, err
		}
		if nil == resp {
			return nil, errNilHTTPResponse
		}
	}

	return dial(resp)
}

func Dial(httprequest *http.Request) (Client, error) {
	if nil == httprequest {
		return nil, errNilHTTPRequest
	}

	var resp *http.Response
	{
		var httpclient http.Client

		var err error

		resp, err = httpclient.Do(httprequest)
		if nil != err {
			return nil, err
		}
		if nil == resp {
			return nil, errNilHTTPResponse
		}
	}

	return dial(resp)
}

func dial(resp *http.Response) (Client, error) {
	if nil == resp {
		return nil, errNilHTTPResponse
	}

	{
		var statuscode int = resp.StatusCode

		if http.StatusOK != statuscode {
			return nil, errhttp.Return(statuscode)
		}
	}

	var body io.ReadCloser = resp.Body
	if nil == body {
		return nil, errNilHTTPResponseBody
	}

	var client = internalClient{
		httpResponse:resp,
		runescanner:utf8.NewRuneScanner(body),
	}

	return &client, nil
}

type internalClient struct {
	httpResponse *http.Response
	runescanner io.RuneScanner
	closed bool
	err error
	nextEvent Event
}

func (receiver *internalClient) Close() error {
	if receiver.closed {
		return nil
	}

	var resp *http.Response = receiver.httpResponse
	if nil == resp {
		return errNilHTTPResponse
	}

	var body io.ReadCloser = resp.Body
	if nil == body {
		return errNilHTTPResponseBody
	}

	err := body.Close()
	if nil != err {
		return err
	}

	receiver.closed = true
	return nil
}

func (receiver internalClient) Decode(dst interface{}) error {

	{
		var err error = receiver.err

		if nil != err {
			return err
		}
	}

	var event *Event
	var casted bool
	{
		event, casted = dst.(*Event)
		if !casted {
			return erorr.Errorf("httpsse: cannot decode into something of type %T", dst)
		}
	}

	*event = receiver.nextEvent
	return nil
}

func (receiver internalClient) Err() error {
	return receiver.err
}

func (receiver *internalClient) Next() bool {

	var runescanner io.RuneScanner = receiver.runescanner
	if nil == runescanner {
		receiver.err = errNilRuneScanner
		return false
	}

	err := readEvent(&receiver.nextEvent, runescanner)
	if nil != err {
		receiver.err = err
		return false
	}

	return true
}
