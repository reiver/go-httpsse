package comment

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
			String:         ":\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{
			String:         ":\r",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{
			String:         ":\r\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},



		{
			String:                   "::\n",
			ExpectedString:            ":",
			ExpectedWritten: int64(len(":")),
		},
		{
			String:                   "::\r",
			ExpectedString:            ":",
			ExpectedWritten: int64(len(":")),
		},
		{
			String:                   "::\r\n",
			ExpectedString:            ":",
			ExpectedWritten: int64(len(":")),
		},



		{
			String:                   ":something\n",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},
		{
			String:                   ":something\r",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},
		{
			String:                   ":something\r\n",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},



		{
			String:                   ": something\n",
			ExpectedString:            " something",
			ExpectedWritten: int64(len(" something")),
		},
		{
			String:                   ": something\r",
			ExpectedString:            " something",
			ExpectedWritten: int64(len(" something")),
		},
		{
			String:                   ": something\r\n",
			ExpectedString:            " something",
			ExpectedWritten: int64(len(" something")),
		},



		{
			String:                   ": Hello world! 🙂\n",
			ExpectedString:            " Hello world! 🙂",
			ExpectedWritten: int64(len(" Hello world! 🙂")),
		},
		{
			String:                   ": Hello world! 🙂\r",
			ExpectedString:            " Hello world! 🙂",
			ExpectedWritten: int64(len(" Hello world! 🙂")),
		},
		{
			String:                   ": Hello world! 🙂\r\n",
			ExpectedString:            " Hello world! 🙂",
			ExpectedWritten: int64(len(" Hello world! 🙂")),
		},



		{
			String:                   ":name:value\n",
			ExpectedString:            "name:value",
			ExpectedWritten: int64(len("name:value")),
		},
		{
			String:                   ":name:value\r",
			ExpectedString:            "name:value",
			ExpectedWritten: int64(len("name:value")),
		},
		{
			String:                   ":name:value\r\n",
			ExpectedString:            "name:value",
			ExpectedWritten: int64(len("name:value")),
		},



		{
			String:                   ": first line\n: second line\n: third line\n",
			ExpectedString:            " first line",
			ExpectedWritten: int64(len(" first line")),
		},
		{
			String:                   ": first line\r: second line\r: third line\r",
			ExpectedString:            " first line",
			ExpectedWritten: int64(len(" first line")),
		},
		{
			String:                   ": first line\r\n: second line\r\n: third line\r\n",
			ExpectedString:            " first line",
			ExpectedWritten: int64(len(" first line")),
		},



		{
			String:                   ":\u0000\r\n",
			ExpectedString:            "\u0000",
			ExpectedWritten: int64(len("\u0000")),
		},
		{
			String:                   ":\u0001\r\n",
			ExpectedString:            "\u0001",
			ExpectedWritten: int64(len("\u0001")),
		},
		{
			String:                   ":\u0002\r\n",
			ExpectedString:            "\u0002",
			ExpectedWritten: int64(len("\u0002")),
		},
		{
			String:                   ":\u0003\r\n",
			ExpectedString:            "\u0003",
			ExpectedWritten: int64(len("\u0003")),
		},
		{
			String:                   ":\u0004\r\n",
			ExpectedString:            "\u0004",
			ExpectedWritten: int64(len("\u0004")),
		},
		{
			String:                   ":\u0005\r\n",
			ExpectedString:            "\u0005",
			ExpectedWritten: int64(len("\u0005")),
		},
		{
			String:                   ":\u0006\r\n",
			ExpectedString:            "\u0006",
			ExpectedWritten: int64(len("\u0006")),
		},
		{
			String:                   ":\u0007\r\n",
			ExpectedString:            "\u0007",
			ExpectedWritten: int64(len("\u0007")),
		},
		{
			String:                   ":\u0008\r\n",
			ExpectedString:            "\u0008",
			ExpectedWritten: int64(len("\u0008")),
		},
		{
			String:                   ":\u0009\r\n",
			ExpectedString:            "\u0009",
			ExpectedWritten: int64(len("\u0009")),
		},



		{
			String:                   ":\u000B\r\n",
			ExpectedString:            "\u000B",
			ExpectedWritten: int64(len("\u000B")),
		},
		{
			String:                   ":\u000C\r\n",
			ExpectedString:            "\u000C",
			ExpectedWritten: int64(len("\u000C")),
		},



		{
			String:                   ":\u000E\r\n",
			ExpectedString:            "\u000E",
			ExpectedWritten: int64(len("\u000E")),
		},
		{
			String:                   ":\u000F\r\n",
			ExpectedString:            "\u000F",
			ExpectedWritten: int64(len("\u000F")),
		},
		{
			String:                   ":\u0010\r\n",
			ExpectedString:            "\u0010",
			ExpectedWritten: int64(len("\u0010")),
		},
		{
			String:                   ":\u0011\r\n",
			ExpectedString:            "\u0011",
			ExpectedWritten: int64(len("\u0011")),
		},
		{
			String:                   ":\u0012\r\n",
			ExpectedString:            "\u0012",
			ExpectedWritten: int64(len("\u0012")),
		},
		{
			String:                   ":\u0013\r\n",
			ExpectedString:            "\u0013",
			ExpectedWritten: int64(len("\u0013")),
		},
		{
			String:                   ":\u0014\r\n",
			ExpectedString:            "\u0014",
			ExpectedWritten: int64(len("\u0014")),
		},
		{
			String:                   ":\u0015\r\n",
			ExpectedString:            "\u0015",
			ExpectedWritten: int64(len("\u0015")),
		},
		{
			String:                   ":\u0016\r\n",
			ExpectedString:            "\u0016",
			ExpectedWritten: int64(len("\u0016")),
		},
		{
			String:                   ":\u0017\r\n",
			ExpectedString:            "\u0017",
			ExpectedWritten: int64(len("\u0017")),
		},
		{
			String:                   ":\u0018\r\n",
			ExpectedString:            "\u0018",
			ExpectedWritten: int64(len("\u0018")),
		},
		{
			String:                   ":\u0019\r\n",
			ExpectedString:            "\u0019",
			ExpectedWritten: int64(len("\u0019")),
		},
		{
			String:                   ":\u001A\r\n",
			ExpectedString:            "\u001A",
			ExpectedWritten: int64(len("\u001A")),
		},
		{
			String:                   ":\u001B\r\n",
			ExpectedString:            "\u001B",
			ExpectedWritten: int64(len("\u001B")),
		},
		{
			String:                   ":\u001C\r\n",
			ExpectedString:            "\u001C",
			ExpectedWritten: int64(len("\u001C")),
		},
		{
			String:                   ":\u001D\r\n",
			ExpectedString:            "\u001D",
			ExpectedWritten: int64(len("\u001D")),
		},
		{
			String:                   ":\u001E\r\n",
			ExpectedString:            "\u001E",
			ExpectedWritten: int64(len("\u001E")),
		},
		{
			String:                   ":\u001F\r\n",
			ExpectedString:            "\u001F",
			ExpectedWritten: int64(len("\u001F")),
		},
		{
			String:                   ":\u0020\r\n",
			ExpectedString:            "\u0020",
			ExpectedWritten: int64(len("\u0020")),
		},
		{
			String:                   ":\u0021\r\n",
			ExpectedString:            "\u0021",
			ExpectedWritten: int64(len("\u0021")),
		},
		{
			String:                   ":\u0022\r\n",
			ExpectedString:            "\u0022",
			ExpectedWritten: int64(len("\u0022")),
		},
		{
			String:                   ":\u0023\r\n",
			ExpectedString:            "\u0023",
			ExpectedWritten: int64(len("\u0023")),
		},
		{
			String:                   ":\u0024\r\n",
			ExpectedString:            "\u0024",
			ExpectedWritten: int64(len("\u0024")),
		},
		{
			String:                   ":\u0025\r\n",
			ExpectedString:            "\u0025",
			ExpectedWritten: int64(len("\u0025")),
		},
		{
			String:                   ":\u0026\r\n",
			ExpectedString:            "\u0026",
			ExpectedWritten: int64(len("\u0026")),
		},
		{
			String:                   ":\u0027\r\n",
			ExpectedString:            "\u0027",
			ExpectedWritten: int64(len("\u0027")),
		},
		{
			String:                   ":\u0028\r\n",
			ExpectedString:            "\u0028",
			ExpectedWritten: int64(len("\u0028")),
		},
		{
			String:                   ":\u0029\r\n",
			ExpectedString:            "\u0029",
			ExpectedWritten: int64(len("\u0029")),
		},
		{
			String:                   ":\u002A\r\n",
			ExpectedString:            "\u002A",
			ExpectedWritten: int64(len("\u002A")),
		},
		{
			String:                   ":\u002B\r\n",
			ExpectedString:            "\u002B",
			ExpectedWritten: int64(len("\u002B")),
		},
		{
			String:                   ":\u002C\r\n",
			ExpectedString:            "\u002C",
			ExpectedWritten: int64(len("\u002C")),
		},
		{
			String:                   ":\u002D\r\n",
			ExpectedString:            "\u002D",
			ExpectedWritten: int64(len("\u002D")),
		},
		{
			String:                   ":\u002E\r\n",
			ExpectedString:            "\u002E",
			ExpectedWritten: int64(len("\u002E")),
		},
		{
			String:                   ":\u002F\r\n",
			ExpectedString:            "\u002F",
			ExpectedWritten: int64(len("\u002F")),
		},
		{
			String:                   ":\u0030\r\n",
			ExpectedString:            "\u0030",
			ExpectedWritten: int64(len("\u0030")),
		},
		{
			String:                   ":\u0031\r\n",
			ExpectedString:            "\u0031",
			ExpectedWritten: int64(len("\u0031")),
		},
		{
			String:                   ":\u0032\r\n",
			ExpectedString:            "\u0032",
			ExpectedWritten: int64(len("\u0032")),
		},
		{
			String:                   ":\u0033\r\n",
			ExpectedString:            "\u0033",
			ExpectedWritten: int64(len("\u0033")),
		},
		{
			String:                   ":\u0034\r\n",
			ExpectedString:            "\u0034",
			ExpectedWritten: int64(len("\u0034")),
		},
		{
			String:                   ":\u0035\r\n",
			ExpectedString:            "\u0035",
			ExpectedWritten: int64(len("\u0035")),
		},
		{
			String:                   ":\u0036\r\n",
			ExpectedString:            "\u0036",
			ExpectedWritten: int64(len("\u0036")),
		},
		{
			String:                   ":\u0037\r\n",
			ExpectedString:            "\u0037",
			ExpectedWritten: int64(len("\u0037")),
		},
		{
			String:                   ":\u0038\r\n",
			ExpectedString:            "\u0038",
			ExpectedWritten: int64(len("\u0038")),
		},
		{
			String:                   ":\u0039\r\n",
			ExpectedString:            "\u0039",
			ExpectedWritten: int64(len("\u0039")),
		},
		{
			String:                   ":\u003A\r\n",
			ExpectedString:            "\u003A",
			ExpectedWritten: int64(len("\u003A")),
		},
		{
			String:                   ":\u003B\r\n",
			ExpectedString:            "\u003B",
			ExpectedWritten: int64(len("\u003B")),
		},
		{
			String:                   ":\u003C\r\n",
			ExpectedString:            "\u003C",
			ExpectedWritten: int64(len("\u003C")),
		},
		{
			String:                   ":\u003D\r\n",
			ExpectedString:            "\u003D",
			ExpectedWritten: int64(len("\u003D")),
		},
		{
			String:                   ":\u003E\r\n",
			ExpectedString:            "\u003E",
			ExpectedWritten: int64(len("\u003E")),
		},
		{
			String:                   ":\u003F\r\n",
			ExpectedString:            "\u003F",
			ExpectedWritten: int64(len("\u003F")),
		},
		{
			String:                   ":\u0040\r\n",
			ExpectedString:            "\u0040",
			ExpectedWritten: int64(len("\u0040")),
		},



		// ...



		{
			String:                   ":\U0010FFFD\r\n",
			ExpectedString:            "\U0010FFFD",
			ExpectedWritten: int64(len("\U0010FFFD")),
		},
		{
			String:                   ":\U0010FFFE\r\n",
			ExpectedString:            "\U0010FFFE",
			ExpectedWritten: int64(len("\U0010FFFE")),
		},
		{
			String:                   ":\U0010FFFF\r\n",
			ExpectedString:            "\U0010FFFF",
			ExpectedWritten: int64(len("\U0010FFFF")),
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
