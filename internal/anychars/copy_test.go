package anychars

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"
)

func TestCopy(t *testing.T) {

	tests := []struct{
		String string
		ExpectedString string
		ExpectedWritten int64
	}{
		{
			String:         "\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{
			String:         "\r",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{
			String:         "\r\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},



		{
			String:         "\u0000\r\n",
			ExpectedString: "\u0000",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0001\r\n",
			ExpectedString: "\u0001",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0002\r\n",
			ExpectedString: "\u0002",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0003\r\n",
			ExpectedString: "\u0003",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0004\r\n",
			ExpectedString: "\u0004",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0005\r\n",
			ExpectedString: "\u0005",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0006\r\n",
			ExpectedString: "\u0006",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0007\r\n",
			ExpectedString: "\u0007",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0008\r\n",
			ExpectedString: "\u0008",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0009\r\n",
			ExpectedString: "\u0009",
			ExpectedWritten: 1,
		},



		{
			String:         "\u000B\r\n",
			ExpectedString: "\u000B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u000C\r\n",
			ExpectedString: "\u000C",
			ExpectedWritten: 1,
		},



		{
			String:         "\u000E\r\n",
			ExpectedString: "\u000E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u000F\r\n",
			ExpectedString: "\u000F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0010\r\n",
			ExpectedString: "\u0010",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0011\r\n",
			ExpectedString: "\u0011",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0012\r\n",
			ExpectedString: "\u0012",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0013\r\n",
			ExpectedString: "\u0013",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0014\r\n",
			ExpectedString: "\u0014",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0015\r\n",
			ExpectedString: "\u0015",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0016\r\n",
			ExpectedString: "\u0016",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0017\r\n",
			ExpectedString: "\u0017",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0018\r\n",
			ExpectedString: "\u0018",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0019\r\n",
			ExpectedString: "\u0019",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001A\r\n",
			ExpectedString: "\u001A",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001B\r\n",
			ExpectedString: "\u001B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001C\r\n",
			ExpectedString: "\u001C",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001D\r\n",
			ExpectedString: "\u001D",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001E\r\n",
			ExpectedString: "\u001E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001F\r\n",
			ExpectedString: "\u001F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0020\r\n",
			ExpectedString: "\u0020",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0021\r\n",
			ExpectedString: "\u0021",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0022\r\n",
			ExpectedString: "\u0022",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0023\r\n",
			ExpectedString: "\u0023",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0024\r\n",
			ExpectedString: "\u0024",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0025\r\n",
			ExpectedString: "\u0025",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0026\r\n",
			ExpectedString: "\u0026",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0027\r\n",
			ExpectedString: "\u0027",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0028\r\n",
			ExpectedString: "\u0028",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0029\r\n",
			ExpectedString: "\u0029",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002A\r\n",
			ExpectedString: "\u002A",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002B\r\n",
			ExpectedString: "\u002B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002C\r\n",
			ExpectedString: "\u002C",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002D\r\n",
			ExpectedString: "\u002D",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002E\r\n",
			ExpectedString: "\u002E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002F\r\n",
			ExpectedString: "\u002F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0030\r\n",
			ExpectedString: "\u0030",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0031\r\n",
			ExpectedString: "\u0031",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0032\r\n",
			ExpectedString: "\u0032",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0033\r\n",
			ExpectedString: "\u0033",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0034\r\n",
			ExpectedString: "\u0034",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0035\r\n",
			ExpectedString: "\u0035",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0036\r\n",
			ExpectedString: "\u0036",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0037\r\n",
			ExpectedString: "\u0037",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0038\r\n",
			ExpectedString: "\u0038",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0039\r\n",
			ExpectedString: "\u0039",
			ExpectedWritten: 1,
		},



		{
			String:         "\u003B\r\n",
			ExpectedString: "\u003B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003C\r\n",
			ExpectedString: "\u003C",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003D\r\n",
			ExpectedString: "\u003D",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003E\r\n",
			ExpectedString: "\u003E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003F\r\n",
			ExpectedString: "\u003F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0040\r\n",
			ExpectedString: "\u0040",
			ExpectedWritten: 1,
		},



		// ...



		{
			String:         "\U0010FFFD\r\n",
			ExpectedString: "\U0010FFFD",
			ExpectedWritten: 4,
		},
		{
			String:         "\U0010FFFE\r\n",
			ExpectedString: "\U0010FFFE",
			ExpectedWritten: 4,
		},
		{
			String:         "\U0010FFFF\r\n",
			ExpectedString: "\U0010FFFF",
			ExpectedWritten: 4,
		},









		{
			String:         "name\r\n",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},



		{
			String:         "name:value\n",
			ExpectedString: "name:value",
			ExpectedWritten: 10,
		},
		{
			String:         "name:value\r",
			ExpectedString: "name:value",
			ExpectedWritten: 10,
		},
		{
			String:         "name:value\r\n",
			ExpectedString: "name:value",
			ExpectedWritten: 10,
		},



		{
			String:         "name: value\n",
			ExpectedString: "name: value",
			ExpectedWritten: 11,
		},
		{
			String:         "name: value\r",
			ExpectedString: "name: value",
			ExpectedWritten: 11,
		},
		{
			String:         "name: value\r\n",
			ExpectedString: "name: value",
			ExpectedWritten: 11,
		},




		{
			String:         "banana: yellow\n",
			ExpectedString: "banana: yellow",
			ExpectedWritten: 14,
		},
		{
			String:         "banana: yellow\r",
			ExpectedString: "banana: yellow",
			ExpectedWritten: 14,
		},
		{
			String:         "banana: yellow\r\n",
			ExpectedString: "banana: yellow",
			ExpectedWritten: 14,
		},



		{
			String:         "Hello world! 🙂: 🎉\n",
			ExpectedString: "Hello world! 🙂: 🎉",
			ExpectedWritten: 23,
		},
		{
			String:         "Hello world! 🙂: 🎉\r",
			ExpectedString: "Hello world! 🙂: 🎉",
			ExpectedWritten: 23,
		},
		{
			String:         "Hello world! 🙂: 🎉\r\n",
			ExpectedString: "Hello world! 🙂: 🎉",
			ExpectedWritten: 23,
		},



		{
			String:         "Hello world! 🙂\u000A",
			ExpectedString: "Hello world! 🙂",
			ExpectedWritten: 17,
		},
		{
			String:         "Hello world! 🙂\u000D",
			ExpectedString: "Hello world! 🙂",
			ExpectedWritten: 17,
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		var reader io.Reader = strings.NewReader(test.String)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualWritten, err := Copy(&actualBuffer, runescanner)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedString
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'string' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedWritten
			actual   :=        actualWritten

			if expected != actual {
				t.Errorf("For test #%d, the actual 'written' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
