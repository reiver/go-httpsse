package field

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"
)

func TestCopy(t *testing.T) {

	tests := []struct{
		String string
		ExpectedName string
		ExpectedValue string
		ExpectedNameWritten int64
		ExpectedValueWritten int64
	}{
		{
			String:                        "name:\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "",
			ExpectedValueWritten:     int64(len("")),
		},
		{
			String:                        "name:\r",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "",
			ExpectedValueWritten:     int64(len("")),
		},
		{
			String:                        "name:\r\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "",
			ExpectedValueWritten:     int64(len("")),
		},



		{
			String:                        "name::\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      ":",
			ExpectedValueWritten:     int64(len(":")),
		},
		{
			String:                        "name::\r",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      ":",
			ExpectedValueWritten:     int64(len(":")),
		},
		{
			String:                        "name::\r\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      ":",
			ExpectedValueWritten:     int64(len(":")),
		},



		{
			String:                        "name:something\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "something",
			ExpectedValueWritten:     int64(len("something")),
		},
		{
			String:                        "name:something\r",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "something",
			ExpectedValueWritten:     int64(len("something")),
		},
		{
			String:                        "name:something\r\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "something",
			ExpectedValueWritten:     int64(len("something")),
		},



		{
			String:                        "name: something\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                       "something",
			ExpectedValueWritten:      int64(len("something")),
		},
		{
			String:                        "name: something\r",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                       "something",
			ExpectedValueWritten:      int64(len("something")),
		},
		{
			String:                        "name: something\r\n",
			ExpectedName:                  "name",
			ExpectedNameWritten: int64(len("name")),
			ExpectedValue:                      "something",
			ExpectedValueWritten:     int64(len("something")),
		},



/*
		{
			String:                   ": Hello world! 🙂\n",
			ExpectedValue:            " Hello world! 🙂",
			ExpectedValueWritten: int64(len(" Hello world! 🙂")),
		},
		{
			String:                   ": Hello world! 🙂\r",
			ExpectedValue:            " Hello world! 🙂",
			ExpectedValueWritten: int64(len(" Hello world! 🙂")),
		},
		{
			String:                   ": Hello world! 🙂\r\n",
			ExpectedValue:            " Hello world! 🙂",
			ExpectedValueWritten: int64(len(" Hello world! 🙂")),
		},



		{
			String:                   ":name:value\n",
			ExpectedValue:            "name:value",
			ExpectedValueWritten: int64(len("name:value")),
		},
		{
			String:                   ":name:value\r",
			ExpectedValue:            "name:value",
			ExpectedValueWritten: int64(len("name:value")),
		},
		{
			String:                   ":name:value\r\n",
			ExpectedValue:            "name:value",
			ExpectedValueWritten: int64(len("name:value")),
		},



		{
			String:                   ": first line\n: second line\n: third line\n",
			ExpectedValue:            " first line",
			ExpectedValueWritten: int64(len(" first line")),
		},
		{
			String:                   ": first line\r: second line\r: third line\r",
			ExpectedValue:            " first line",
			ExpectedValueWritten: int64(len(" first line")),
		},
		{
			String:                   ": first line\r\n: second line\r\n: third line\r\n",
			ExpectedValue:            " first line",
			ExpectedValueWritten: int64(len(" first line")),
		},



		{
			String:                   ":\u0000\r\n",
			ExpectedValue:            "\u0000",
			ExpectedValueWritten: int64(len("\u0000")),
		},
		{
			String:                   ":\u0001\r\n",
			ExpectedValue:            "\u0001",
			ExpectedValueWritten: int64(len("\u0001")),
		},
		{
			String:                   ":\u0002\r\n",
			ExpectedValue:            "\u0002",
			ExpectedValueWritten: int64(len("\u0002")),
		},
		{
			String:                   ":\u0003\r\n",
			ExpectedValue:            "\u0003",
			ExpectedValueWritten: int64(len("\u0003")),
		},
		{
			String:                   ":\u0004\r\n",
			ExpectedValue:            "\u0004",
			ExpectedValueWritten: int64(len("\u0004")),
		},
		{
			String:                   ":\u0005\r\n",
			ExpectedValue:            "\u0005",
			ExpectedValueWritten: int64(len("\u0005")),
		},
		{
			String:                   ":\u0006\r\n",
			ExpectedValue:            "\u0006",
			ExpectedValueWritten: int64(len("\u0006")),
		},
		{
			String:                   ":\u0007\r\n",
			ExpectedValue:            "\u0007",
			ExpectedValueWritten: int64(len("\u0007")),
		},
		{
			String:                   ":\u0008\r\n",
			ExpectedValue:            "\u0008",
			ExpectedValueWritten: int64(len("\u0008")),
		},
		{
			String:                   ":\u0009\r\n",
			ExpectedValue:            "\u0009",
			ExpectedValueWritten: int64(len("\u0009")),
		},



		{
			String:                   ":\u000B\r\n",
			ExpectedValue:            "\u000B",
			ExpectedValueWritten: int64(len("\u000B")),
		},
		{
			String:                   ":\u000C\r\n",
			ExpectedValue:            "\u000C",
			ExpectedValueWritten: int64(len("\u000C")),
		},



		{
			String:                   ":\u000E\r\n",
			ExpectedValue:            "\u000E",
			ExpectedValueWritten: int64(len("\u000E")),
		},
		{
			String:                   ":\u000F\r\n",
			ExpectedValue:            "\u000F",
			ExpectedValueWritten: int64(len("\u000F")),
		},
		{
			String:                   ":\u0010\r\n",
			ExpectedValue:            "\u0010",
			ExpectedValueWritten: int64(len("\u0010")),
		},
		{
			String:                   ":\u0011\r\n",
			ExpectedValue:            "\u0011",
			ExpectedValueWritten: int64(len("\u0011")),
		},
		{
			String:                   ":\u0012\r\n",
			ExpectedValue:            "\u0012",
			ExpectedValueWritten: int64(len("\u0012")),
		},
		{
			String:                   ":\u0013\r\n",
			ExpectedValue:            "\u0013",
			ExpectedValueWritten: int64(len("\u0013")),
		},
		{
			String:                   ":\u0014\r\n",
			ExpectedValue:            "\u0014",
			ExpectedValueWritten: int64(len("\u0014")),
		},
		{
			String:                   ":\u0015\r\n",
			ExpectedValue:            "\u0015",
			ExpectedValueWritten: int64(len("\u0015")),
		},
		{
			String:                   ":\u0016\r\n",
			ExpectedValue:            "\u0016",
			ExpectedValueWritten: int64(len("\u0016")),
		},
		{
			String:                   ":\u0017\r\n",
			ExpectedValue:            "\u0017",
			ExpectedValueWritten: int64(len("\u0017")),
		},
		{
			String:                   ":\u0018\r\n",
			ExpectedValue:            "\u0018",
			ExpectedValueWritten: int64(len("\u0018")),
		},
		{
			String:                   ":\u0019\r\n",
			ExpectedValue:            "\u0019",
			ExpectedValueWritten: int64(len("\u0019")),
		},
		{
			String:                   ":\u001A\r\n",
			ExpectedValue:            "\u001A",
			ExpectedValueWritten: int64(len("\u001A")),
		},
		{
			String:                   ":\u001B\r\n",
			ExpectedValue:            "\u001B",
			ExpectedValueWritten: int64(len("\u001B")),
		},
		{
			String:                   ":\u001C\r\n",
			ExpectedValue:            "\u001C",
			ExpectedValueWritten: int64(len("\u001C")),
		},
		{
			String:                   ":\u001D\r\n",
			ExpectedValue:            "\u001D",
			ExpectedValueWritten: int64(len("\u001D")),
		},
		{
			String:                   ":\u001E\r\n",
			ExpectedValue:            "\u001E",
			ExpectedValueWritten: int64(len("\u001E")),
		},
		{
			String:                   ":\u001F\r\n",
			ExpectedValue:            "\u001F",
			ExpectedValueWritten: int64(len("\u001F")),
		},
		{
			String:                   ":\u0020\r\n",
			ExpectedValue:            "\u0020",
			ExpectedValueWritten: int64(len("\u0020")),
		},
		{
			String:                   ":\u0021\r\n",
			ExpectedValue:            "\u0021",
			ExpectedValueWritten: int64(len("\u0021")),
		},
		{
			String:                   ":\u0022\r\n",
			ExpectedValue:            "\u0022",
			ExpectedValueWritten: int64(len("\u0022")),
		},
		{
			String:                   ":\u0023\r\n",
			ExpectedValue:            "\u0023",
			ExpectedValueWritten: int64(len("\u0023")),
		},
		{
			String:                   ":\u0024\r\n",
			ExpectedValue:            "\u0024",
			ExpectedValueWritten: int64(len("\u0024")),
		},
		{
			String:                   ":\u0025\r\n",
			ExpectedValue:            "\u0025",
			ExpectedValueWritten: int64(len("\u0025")),
		},
		{
			String:                   ":\u0026\r\n",
			ExpectedValue:            "\u0026",
			ExpectedValueWritten: int64(len("\u0026")),
		},
		{
			String:                   ":\u0027\r\n",
			ExpectedValue:            "\u0027",
			ExpectedValueWritten: int64(len("\u0027")),
		},
		{
			String:                   ":\u0028\r\n",
			ExpectedValue:            "\u0028",
			ExpectedValueWritten: int64(len("\u0028")),
		},
		{
			String:                   ":\u0029\r\n",
			ExpectedValue:            "\u0029",
			ExpectedValueWritten: int64(len("\u0029")),
		},
		{
			String:                   ":\u002A\r\n",
			ExpectedValue:            "\u002A",
			ExpectedValueWritten: int64(len("\u002A")),
		},
		{
			String:                   ":\u002B\r\n",
			ExpectedValue:            "\u002B",
			ExpectedValueWritten: int64(len("\u002B")),
		},
		{
			String:                   ":\u002C\r\n",
			ExpectedValue:            "\u002C",
			ExpectedValueWritten: int64(len("\u002C")),
		},
		{
			String:                   ":\u002D\r\n",
			ExpectedValue:            "\u002D",
			ExpectedValueWritten: int64(len("\u002D")),
		},
		{
			String:                   ":\u002E\r\n",
			ExpectedValue:            "\u002E",
			ExpectedValueWritten: int64(len("\u002E")),
		},
		{
			String:                   ":\u002F\r\n",
			ExpectedValue:            "\u002F",
			ExpectedValueWritten: int64(len("\u002F")),
		},
		{
			String:                   ":\u0030\r\n",
			ExpectedValue:            "\u0030",
			ExpectedValueWritten: int64(len("\u0030")),
		},
		{
			String:                   ":\u0031\r\n",
			ExpectedValue:            "\u0031",
			ExpectedValueWritten: int64(len("\u0031")),
		},
		{
			String:                   ":\u0032\r\n",
			ExpectedValue:            "\u0032",
			ExpectedValueWritten: int64(len("\u0032")),
		},
		{
			String:                   ":\u0033\r\n",
			ExpectedValue:            "\u0033",
			ExpectedValueWritten: int64(len("\u0033")),
		},
		{
			String:                   ":\u0034\r\n",
			ExpectedValue:            "\u0034",
			ExpectedValueWritten: int64(len("\u0034")),
		},
		{
			String:                   ":\u0035\r\n",
			ExpectedValue:            "\u0035",
			ExpectedValueWritten: int64(len("\u0035")),
		},
		{
			String:                   ":\u0036\r\n",
			ExpectedValue:            "\u0036",
			ExpectedValueWritten: int64(len("\u0036")),
		},
		{
			String:                   ":\u0037\r\n",
			ExpectedValue:            "\u0037",
			ExpectedValueWritten: int64(len("\u0037")),
		},
		{
			String:                   ":\u0038\r\n",
			ExpectedValue:            "\u0038",
			ExpectedValueWritten: int64(len("\u0038")),
		},
		{
			String:                   ":\u0039\r\n",
			ExpectedValue:            "\u0039",
			ExpectedValueWritten: int64(len("\u0039")),
		},
		{
			String:                   ":\u003A\r\n",
			ExpectedValue:            "\u003A",
			ExpectedValueWritten: int64(len("\u003A")),
		},
		{
			String:                   ":\u003B\r\n",
			ExpectedValue:            "\u003B",
			ExpectedValueWritten: int64(len("\u003B")),
		},
		{
			String:                   ":\u003C\r\n",
			ExpectedValue:            "\u003C",
			ExpectedValueWritten: int64(len("\u003C")),
		},
		{
			String:                   ":\u003D\r\n",
			ExpectedValue:            "\u003D",
			ExpectedValueWritten: int64(len("\u003D")),
		},
		{
			String:                   ":\u003E\r\n",
			ExpectedValue:            "\u003E",
			ExpectedValueWritten: int64(len("\u003E")),
		},
		{
			String:                   ":\u003F\r\n",
			ExpectedValue:            "\u003F",
			ExpectedValueWritten: int64(len("\u003F")),
		},
		{
			String:                   ":\u0040\r\n",
			ExpectedValue:            "\u0040",
			ExpectedValueWritten: int64(len("\u0040")),
		},



		// ...



		{
			String:                   ":\U0010FFFD\r\n",
			ExpectedValue:            "\U0010FFFD",
			ExpectedValueWritten: int64(len("\U0010FFFD")),
		},
		{
			String:                   ":\U0010FFFE\r\n",
			ExpectedValue:            "\U0010FFFE",
			ExpectedValueWritten: int64(len("\U0010FFFE")),
		},
		{
			String:                   ":\U0010FFFF\r\n",
			ExpectedValue:            "\U0010FFFF",
			ExpectedValueWritten: int64(len("\U0010FFFF")),
		},
*/
	}

	for testNumber, test := range tests {

		var actualNameBuffer  strings.Builder
		var actualValueBuffer strings.Builder

		var reader io.Reader = strings.NewReader(test.String)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualNameWritten, actualValueWritten, err := Copy(&actualNameBuffer, &actualValueBuffer, runescanner)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedName
			actual   := actualNameBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := actualValueBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'value' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedNameWritten
			actual   :=        actualNameWritten

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name-written' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedValueWritten
			actual   :=        actualValueWritten

			if expected != actual {
				t.Errorf("For test #%d, the actual 'value-written' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
