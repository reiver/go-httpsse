package httpsse

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"
)

func TestReadEvent(t *testing.T) {

	tests := []struct{
		String string
		ExpectedEventName string
		ExpectedEventID string
		ExpectedEventData string
	}{
		{ // 0
			String:
				"event:banana"+"\n"+
				""            +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{ // 1
			String:
				"event:banana"+"\r"+
				""            +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{ // 2
			String:
				"event:banana"+"\r\n"+
				""            +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},



		{ // 3
			String:
				"event: banana"+"\n"+
				""             +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{ // 4
			String:
				"event: banana"+"\r"+
				""             +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{ // 5
			String:
				"event: banana"+"\r\n"+
				""             +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},









		{ // 6
			String:
				"event:banana"+"\n"+
				"id:123"      +"\n"+
				""            +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{ // 7
			String:
				"event:banana"+"\r"+
				"id:123"      +"\r"+
				""            +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{ // 8
			String:
				"event:banana"+"\r\n"+
				"id:123"      +"\r\n"+
				""            +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},



		{ // 9
			String:
				"event: banana"+"\n"+
				"id: 123"      +"\n"+
				""             +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{ // 10
			String:
				"event: banana"+"\r"+
				"id: 123"      +"\r"+
				""             +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{ // 11
			String:
				"event: banana"+"\r\n"+
				"id: 123"      +"\r\n"+
				""             +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},









		{ // 12
			String:
				"id:123"      +"\n"+
				""            +"\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{ // 13
			String:
				"id:123"      +"\r"+
				""            +"\r",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{ // 14
			String:
				"id:123"      +"\r\n"+
				""            +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},



		{ // 15
			String:
				"id: 123"      +"\n"+
				""             +"\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{ // 16
			String:
				"id: 123"      +"\r"+
				""             +"\r",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{ // 17
			String:
				"id: 123"      +"\r\n"+
				""             +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},









		{ // 18
			String:
				": I was here"    +"\r\n"+
				"event: something"+"\r\n"+
				"id: 0xABCD"      +"\r\n"+
				""                +"\r\n",
			ExpectedEventName:"something",
			ExpectedEventID:"0xABCD",
		},









		{ // 19
			String:
				"data: once"   +"\r\n"+
				"data: twice"  +"\r\n"+
				"data: thrice" +"\r\n"+
				"data: fource" +"\r\n"+
				""          +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:
				"once"   +"\n"+
				"twice"  +"\n"+
				"thrice" +"\n"+
				"fource",
		},



		{ // 20
			String:
				"id: once"   +"\r\n"+
				"id: twice"  +"\r\n"+
				"id: thrice" +"\r\n"+
				"id: fource" +"\r\n"+
				""          +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"fource",
			ExpectedEventData:"",
		},



		{ // 21
			String:
				"event: once"   +"\r\n"+
				"event: twice"  +"\r\n"+
				"event: thrice" +"\r\n"+
				"event: fource" +"\r\n"+
				""          +"\r\n",
			ExpectedEventName:"fource",
			ExpectedEventID:"",
			ExpectedEventData:"",
		},









		{ // 22
			String:
				"data: YHOO" +"\r\n"+
				"data: +2"   +"\r\n"+
				"data: 10"   +"\r\n"+
				""          +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"YHOO\n+2\n10",
		},









		{ // 23
			String:
				": test stream"      +"\r\n"+
				""                   +"\r\n"+
				"data: first event"  +"\r\n"+
				"id: 1"              +"\r\n"+
				""                   +"\r\n"+
				"data:second event"  +"\r\n"+
				"id"                 +"\r\n"+
				""                   +"\r\n"+
				"data:  third event" +"\r\n"+
				""                   +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"",
		},
		{ // 24
			String:
				"data: first event"  +"\r\n"+
				"id: 1"              +"\r\n"+
				""                   +"\r\n"+
				"data:second event"  +"\r\n"+
				"id"                 +"\r\n"+
				""                   +"\r\n"+
				"data:  third event" +"\r\n"+
				""                   +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"1",
			ExpectedEventData:"first event",
		},
		{ // 25
			String:
				"data:second event"  +"\r\n"+
				"id"                 +"\r\n"+
				""                   +"\r\n"+
				"data:  third event" +"\r\n"+
				""                   +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"second event",
		},
		{ // 26
			String:
				"data:  third event" +"\r\n"+
				""                   +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:" third event",
		},









		{ // 27
			String:
				"data"  +"\r\n"+
				""      +"\r\n"+
				"data"  +"\r\n"+
				"data"  +"\r\n"+
				""      +"\r\n"+
				"data:" +"\r\n"+
				""      +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"",
		},
		{ // 28
			String:
				"data"  +"\r\n"+
				"data"  +"\r\n"+
				""      +"\r\n"+
				"data:" +"\r\n"+
				""      +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"\n",
		},
		{ // 29
			String:
				"data:" +"\r\n"+
				""      +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"",
		},









		{ // 30
			String:
				"data:test"  +"\r\n"+
				""           +"\r\n"+
				"data: test" +"\r\n"+
				""           +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"test",
		},
		{ // 31
			String:
				"data: test" +"\r\n"+
				""           +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"test",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.String)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		var event Event

		err := readEvent(&event, runescanner)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			t.Logf("EVENT:\n%v", event)
			continue
		}

		{
			expected := test.ExpectedEventName
			actual := event.EventName()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'event-name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				t.Logf("EVENT:\n%v", event)
				continue
			}
		}

		{
			expected := test.ExpectedEventID
			actual := event.EventID()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'event-id' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				t.Logf("EVENT:\n%v", event)
				continue
			}
		}

		{
			expected := test.ExpectedEventData
			actual := event.EventDataCollapsed()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'event-data' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				t.Logf("EVENT:\n%v", event)
				continue
			}
		}
	}
}
