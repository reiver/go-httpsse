package name

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
			String:         "\u0000",
			ExpectedString: "\u0000",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0001",
			ExpectedString: "\u0001",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0002",
			ExpectedString: "\u0002",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0003",
			ExpectedString: "\u0003",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0004",
			ExpectedString: "\u0004",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0005",
			ExpectedString: "\u0005",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0006",
			ExpectedString: "\u0006",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0007",
			ExpectedString: "\u0007",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0008",
			ExpectedString: "\u0008",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0009",
			ExpectedString: "\u0009",
			ExpectedWritten: 1,
		},



		{
			String:         "\u000B",
			ExpectedString: "\u000B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u000C",
			ExpectedString: "\u000C",
			ExpectedWritten: 1,
		},



		{
			String:         "\u000E",
			ExpectedString: "\u000E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u000F",
			ExpectedString: "\u000F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0010",
			ExpectedString: "\u0010",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0011",
			ExpectedString: "\u0011",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0012",
			ExpectedString: "\u0012",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0013",
			ExpectedString: "\u0013",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0014",
			ExpectedString: "\u0014",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0015",
			ExpectedString: "\u0015",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0016",
			ExpectedString: "\u0016",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0017",
			ExpectedString: "\u0017",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0018",
			ExpectedString: "\u0018",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0019",
			ExpectedString: "\u0019",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001A",
			ExpectedString: "\u001A",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001B",
			ExpectedString: "\u001B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001C",
			ExpectedString: "\u001C",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001D",
			ExpectedString: "\u001D",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001E",
			ExpectedString: "\u001E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u001F",
			ExpectedString: "\u001F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0020",
			ExpectedString: "\u0020",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0021",
			ExpectedString: "\u0021",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0022",
			ExpectedString: "\u0022",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0023",
			ExpectedString: "\u0023",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0024",
			ExpectedString: "\u0024",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0025",
			ExpectedString: "\u0025",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0026",
			ExpectedString: "\u0026",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0027",
			ExpectedString: "\u0027",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0028",
			ExpectedString: "\u0028",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0029",
			ExpectedString: "\u0029",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002A",
			ExpectedString: "\u002A",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002B",
			ExpectedString: "\u002B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002C",
			ExpectedString: "\u002C",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002D",
			ExpectedString: "\u002D",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002E",
			ExpectedString: "\u002E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u002F",
			ExpectedString: "\u002F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0030",
			ExpectedString: "\u0030",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0031",
			ExpectedString: "\u0031",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0032",
			ExpectedString: "\u0032",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0033",
			ExpectedString: "\u0033",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0034",
			ExpectedString: "\u0034",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0035",
			ExpectedString: "\u0035",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0036",
			ExpectedString: "\u0036",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0037",
			ExpectedString: "\u0037",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0038",
			ExpectedString: "\u0038",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0039",
			ExpectedString: "\u0039",
			ExpectedWritten: 1,
		},



		{
			String:         "\u003B",
			ExpectedString: "\u003B",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003C",
			ExpectedString: "\u003C",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003D",
			ExpectedString: "\u003D",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003E",
			ExpectedString: "\u003E",
			ExpectedWritten: 1,
		},
		{
			String:         "\u003F",
			ExpectedString: "\u003F",
			ExpectedWritten: 1,
		},
		{
			String:         "\u0040",
			ExpectedString: "\u0040",
			ExpectedWritten: 1,
		},



		// ...



		{
			String:         "\U0010FFFD",
			ExpectedString: "\U0010FFFD",
			ExpectedWritten: 4,
		},
		{
			String:         "\U0010FFFE",
			ExpectedString: "\U0010FFFE",
			ExpectedWritten: 4,
		},
		{
			String:         "\U0010FFFF",
			ExpectedString: "\U0010FFFF",
			ExpectedWritten: 4,
		},









		{
			String:         "name:value",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},
		{
			String:         "name:value\n",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},
		{
			String:         "name:value\r",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},
		{
			String:         "name:value\r\n",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},



		{
			String:         "name: value",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},
		{
			String:         "name: value\n",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},
		{
			String:         "name: value\r",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},
		{
			String:         "name: value\r\n",
			ExpectedString: "name",
			ExpectedWritten: 4,
		},




		{
			String:         "banana: yellow",
			ExpectedString: "banana",
			ExpectedWritten: 6,
		},
		{
			String:         "banana: yellow\n",
			ExpectedString: "banana",
			ExpectedWritten: 6,
		},
		{
			String:         "banana: yellow\r",
			ExpectedString: "banana",
			ExpectedWritten: 6,
		},
		{
			String:         "banana: yellow\r\n",
			ExpectedString: "banana",
			ExpectedWritten: 6,
		},



		{
			String:         "Hello world! 🙂: 🎉",
			ExpectedString: "Hello world! 🙂",
			ExpectedWritten: 17,
		},
		{
			String:         "Hello world! 🙂: 🎉\n",
			ExpectedString: "Hello world! 🙂",
			ExpectedWritten: 17,
		},
		{
			String:         "Hello world! 🙂: 🎉\r",
			ExpectedString: "Hello world! 🙂",
			ExpectedWritten: 17,
		},
		{
			String:         "Hello world! 🙂: 🎉\r\n",
			ExpectedString: "Hello world! 🙂",
			ExpectedWritten: 17,
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
		{
			String:         "Hello world! 🙂\u003A",
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
			continue
		}

		{
			expected := test.ExpectedString
			actual   := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'string' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
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
				continue
			}
		}
	}
}
