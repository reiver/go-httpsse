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
		{
			String:
				"event:banana"+"\n"+
				""            +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{
			String:
				"event:banana"+"\r"+
				""            +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{
			String:
				"event:banana"+"\r\n"+
				""            +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},



		{
			String:
				"event: banana"+"\n"+
				""             +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{
			String:
				"event: banana"+"\r"+
				""             +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},
		{
			String:
				"event: banana"+"\r\n"+
				""             +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"",
		},









		{
			String:
				"event:banana"+"\n"+
				"id:123"      +"\n"+
				""            +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{
			String:
				"event:banana"+"\r"+
				"id:123"      +"\r"+
				""            +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{
			String:
				"event:banana"+"\r\n"+
				"id:123"      +"\r\n"+
				""            +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},



		{
			String:
				"event: banana"+"\n"+
				"id: 123"      +"\n"+
				""             +"\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{
			String:
				"event: banana"+"\r"+
				"id: 123"      +"\r"+
				""             +"\r",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},
		{
			String:
				"event: banana"+"\r\n"+
				"id: 123"      +"\r\n"+
				""             +"\r\n",
			ExpectedEventName:"banana",
			ExpectedEventID:"123",
		},









		{
			String:
				"id:123"      +"\n"+
				""            +"\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{
			String:
				"id:123"      +"\r"+
				""            +"\r",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{
			String:
				"id:123"      +"\r\n"+
				""            +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},



		{
			String:
				"id: 123"      +"\n"+
				""             +"\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{
			String:
				"id: 123"      +"\r"+
				""             +"\r",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},
		{
			String:
				"id: 123"      +"\r\n"+
				""             +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"123",
		},









		{
			String:
				": I was here"    +"\r\n"+
				"event: something"+"\r\n"+
				"id: 0xABCD"      +"\r\n"+
				""                +"\r\n",
			ExpectedEventName:"something",
			ExpectedEventID:"0xABCD",
		},









		{
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



		{
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



		{
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









		{
			String:
				"data: YHOO" +"\r\n"+
				"data: +2"   +"\r\n"+
				"data: 10"   +"\r\n"+
				""          +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"YHOO\n+2\n10",
		},









		{
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
		{
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
		{
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
		{
			String:
				"data:  third event" +"\r\n"+
				""                   +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:" third event",
		},









		{
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
		{
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
		{
			String:
				"data:" +"\r\n"+
				""      +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"",
		},









		{
			String:
				"data:test"  +"\r\n"+
				""           +"\r\n"+
				"data: test" +"\r\n"+
				""           +"\r\n",
			ExpectedEventName:"",
			ExpectedEventID:"",
			ExpectedEventData:"test",
		},
		{
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
			actual := event.EventData()

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
