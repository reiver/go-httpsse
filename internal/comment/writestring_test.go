package comment

import (
	"testing"

	"strings"
)

func TestWriteString(t *testing.T) {

	tests := []struct{
		String string
		Expected string
	}{
		{
			String:    "",
			Expected: ":\n",
		},



		{
			String:    " ",
			Expected: ": \n",
		},
		{
			String:    "  ",
			Expected: ":  \n",
		},
		{
			String:    " ",
			Expected: ": \n",
		},
		{
			String:    "   ",
			Expected: ":   \n",
		},



		{
			String:    ":",
			Expected: "::\n",
		},
		{
			String:    "::",
			Expected: ":::\n",
		},
		{
			String:    ":::",
			Expected: "::::\n",
		},
		{
			String:    "::::",
			Expected: ":::::\n",
		},



		{
			String:    "ping",
			Expected: ":ping\n",
		},
		{
			String:    " ping",
			Expected: ": ping\n",
		},



		{
			String:    "ping\npong",
			Expected: ":ping\n:pong\n",
		},
		{
			String:    "ping\rpong",
			Expected: ":ping\n:pong\n",
		},
		{
			String:    "ping\r\npong",
			Expected: ":ping\n:pong\n",
		},



		{
			String:    "bing\nbong\nbang",
			Expected: ":bing\n:bong\n:bang\n",
		},
		{
			String:    "bing\rbong\rbang",
			Expected: ":bing\n:bong\n:bang\n",
		},
		{
			String:    "bing\r\nbong\r\nbang",
			Expected: ":bing\n:bong\n:bang\n",
		},



		{
			String:    "bing\nbong\rbang\r\nit's",
			Expected: ":bing\n:bong\n:bang\n:it's\n",
		},
	}

	for testNumber, test := range tests {

		var actualBuilder strings.Builder

		_, err := WriteString(&actualBuilder, test.String)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.Expected
			actual := actualBuilder.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual written 'comment' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
