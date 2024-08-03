package fieldvalue

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
		{ // 0
			String:          "\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{ // 1
			String:          "\r",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{ // 2
			String:         "\r\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},



		{ // 3
			String:        ":\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{ // 4
			String:        ":\r",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{ // 5
			String:        ":\r\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},



		{ // 6
			String:        ": \n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{ // 7
			String:        ": \r",
			ExpectedString: "",
			ExpectedWritten: 0,
		},
		{ // 8
			String:        ": \r\n",
			ExpectedString: "",
			ExpectedWritten: 0,
		},



		{ // 9
			String:                   "::\n",
			ExpectedString:            ":",
			ExpectedWritten: int64(len(":")),
		},
		{ // 10
			String:                   "::\r",
			ExpectedString:            ":",
			ExpectedWritten: int64(len(":")),
		},
		{ // 11
			String:                   "::\r\n",
			ExpectedString:            ":",
			ExpectedWritten: int64(len(":")),
		},



		{ // 12
			String:                   ":something\n",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},
		{ // 13
			String:                   ":something\r",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},
		{ // 14
			String:                   ":something\r\n",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},



		{ // 15
			String:                  ": something\n",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},
		{ // 16
			String:                  ": something\r",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},
		{ // 17
			String:                  ": something\r\n",
			ExpectedString:            "something",
			ExpectedWritten: int64(len("something")),
		},



		{ // 18
			String:                   ":Hello world! 🙂\n",
			ExpectedString:            "Hello world! 🙂",
			ExpectedWritten: int64(len("Hello world! 🙂")),
		},
		{ // 19
			String:                   ":Hello world! 🙂\r",
			ExpectedString:            "Hello world! 🙂",
			ExpectedWritten: int64(len("Hello world! 🙂")),
		},
		{ // 20
			String:                   ":Hello world! 🙂\r\n",
			ExpectedString:            "Hello world! 🙂",
			ExpectedWritten: int64(len("Hello world! 🙂")),
		},



		{ // 21
			String:                  ": Hello world! 🙂\n",
			ExpectedString:            "Hello world! 🙂",
			ExpectedWritten: int64(len("Hello world! 🙂")),
		},
		{ // 22
			String:                  ": Hello world! 🙂\r",
			ExpectedString:            "Hello world! 🙂",
			ExpectedWritten: int64(len("Hello world! 🙂")),
		},
		{ // 23
			String:                  ": Hello world! 🙂\r\n",
			ExpectedString:            "Hello world! 🙂",
			ExpectedWritten: int64(len("Hello world! 🙂")),
		},



		{ // 24
			String:                   ":\u0000\r\n",
			ExpectedString:            "\u0000",
			ExpectedWritten: int64(len("\u0000")),
		},
		{ // 25
			String:                   ":\u0001\r\n",
			ExpectedString:            "\u0001",
			ExpectedWritten: int64(len("\u0001")),
		},
		{ // 26
			String:                   ":\u0002\r\n",
			ExpectedString:            "\u0002",
			ExpectedWritten: int64(len("\u0002")),
		},
		{ // 27
			String:                   ":\u0003\r\n",
			ExpectedString:            "\u0003",
			ExpectedWritten: int64(len("\u0003")),
		},
		{ // 28
			String:                   ":\u0004\r\n",
			ExpectedString:            "\u0004",
			ExpectedWritten: int64(len("\u0004")),
		},
		{ // 29
			String:                   ":\u0005\r\n",
			ExpectedString:            "\u0005",
			ExpectedWritten: int64(len("\u0005")),
		},
		{ // 30
			String:                   ":\u0006\r\n",
			ExpectedString:            "\u0006",
			ExpectedWritten: int64(len("\u0006")),
		},
		{ // 31
			String:                   ":\u0007\r\n",
			ExpectedString:            "\u0007",
			ExpectedWritten: int64(len("\u0007")),
		},
		{ // 32
			String:                   ":\u0008\r\n",
			ExpectedString:            "\u0008",
			ExpectedWritten: int64(len("\u0008")),
		},
		{ // 33
			String:                   ":\u0009\r\n",
			ExpectedString:            "\u0009",
			ExpectedWritten: int64(len("\u0009")),
		},



		{ // 34
			String:                   ":\u000B\r\n",
			ExpectedString:            "\u000B",
			ExpectedWritten: int64(len("\u000B")),
		},
		{ // 35
			String:                   ":\u000C\r\n",
			ExpectedString:            "\u000C",
			ExpectedWritten: int64(len("\u000C")),
		},



		{ // 36
			String:                   ":\u000E\r\n",
			ExpectedString:            "\u000E",
			ExpectedWritten: int64(len("\u000E")),
		},
		{ // 37
			String:                   ":\u000F\r\n",
			ExpectedString:            "\u000F",
			ExpectedWritten: int64(len("\u000F")),
		},
		{ // 38
			String:                   ":\u0010\r\n",
			ExpectedString:            "\u0010",
			ExpectedWritten: int64(len("\u0010")),
		},
		{ // 39
			String:                   ":\u0011\r\n",
			ExpectedString:            "\u0011",
			ExpectedWritten: int64(len("\u0011")),
		},
		{ // 40
			String:                   ":\u0012\r\n",
			ExpectedString:            "\u0012",
			ExpectedWritten: int64(len("\u0012")),
		},
		{ // 41
			String:                   ":\u0013\r\n",
			ExpectedString:            "\u0013",
			ExpectedWritten: int64(len("\u0013")),
		},
		{ // 42
			String:                   ":\u0014\r\n",
			ExpectedString:            "\u0014",
			ExpectedWritten: int64(len("\u0014")),
		},
		{ // 43
			String:                   ":\u0015\r\n",
			ExpectedString:            "\u0015",
			ExpectedWritten: int64(len("\u0015")),
		},
		{ // 44
			String:                   ":\u0016\r\n",
			ExpectedString:            "\u0016",
			ExpectedWritten: int64(len("\u0016")),
		},
		{ // 45
			String:                   ":\u0017\r\n",
			ExpectedString:            "\u0017",
			ExpectedWritten: int64(len("\u0017")),
		},
		{ // 46
			String:                   ":\u0018\r\n",
			ExpectedString:            "\u0018",
			ExpectedWritten: int64(len("\u0018")),
		},
		{ // 47
			String:                   ":\u0019\r\n",
			ExpectedString:            "\u0019",
			ExpectedWritten: int64(len("\u0019")),
		},
		{ // 48
			String:                   ":\u001A\r\n",
			ExpectedString:            "\u001A",
			ExpectedWritten: int64(len("\u001A")),
		},
		{ // 49
			String:                   ":\u001B\r\n",
			ExpectedString:            "\u001B",
			ExpectedWritten: int64(len("\u001B")),
		},
		{ // 50
			String:                   ":\u001C\r\n",
			ExpectedString:            "\u001C",
			ExpectedWritten: int64(len("\u001C")),
		},
		{ // 51
			String:                   ":\u001D\r\n",
			ExpectedString:            "\u001D",
			ExpectedWritten: int64(len("\u001D")),
		},
		{ // 52
			String:                   ":\u001E\r\n",
			ExpectedString:            "\u001E",
			ExpectedWritten: int64(len("\u001E")),
		},
		{ // 53
			String:                   ":\u001F\r\n",
			ExpectedString:            "\u001F",
			ExpectedWritten: int64(len("\u001F")),
		},
		{ // 54
			String:                   ":\u0020\r\n", // SP
			ExpectedString:            "" ,
			ExpectedWritten: int64(len("")),
		},
		{ // 55
			String:                   ":\u0021\r\n",
			ExpectedString:            "\u0021",
			ExpectedWritten: int64(len("\u0021")),
		}, // 56
		{
			String:                   ":\u0022\r\n",
			ExpectedString:            "\u0022",
			ExpectedWritten: int64(len("\u0022")),
		},
		{ // 57
			String:                   ":\u0023\r\n",
			ExpectedString:            "\u0023",
			ExpectedWritten: int64(len("\u0023")),
		},
		{ // 58
			String:                   ":\u0024\r\n",
			ExpectedString:            "\u0024",
			ExpectedWritten: int64(len("\u0024")),
		},
		{ // 59
			String:                   ":\u0025\r\n",
			ExpectedString:            "\u0025",
			ExpectedWritten: int64(len("\u0025")),
		},
		{ // 60
			String:                   ":\u0026\r\n",
			ExpectedString:            "\u0026",
			ExpectedWritten: int64(len("\u0026")),
		},
		{ // 61
			String:                   ":\u0027\r\n",
			ExpectedString:            "\u0027",
			ExpectedWritten: int64(len("\u0027")),
		},
		{ // 62
			String:                   ":\u0028\r\n",
			ExpectedString:            "\u0028",
			ExpectedWritten: int64(len("\u0028")),
		},
		{ // 63
			String:                   ":\u0029\r\n",
			ExpectedString:            "\u0029",
			ExpectedWritten: int64(len("\u0029")),
		},
		{ // 64
			String:                   ":\u002A\r\n",
			ExpectedString:            "\u002A",
			ExpectedWritten: int64(len("\u002A")),
		},
		{ // 65
			String:                   ":\u002B\r\n",
			ExpectedString:            "\u002B",
			ExpectedWritten: int64(len("\u002B")),
		},
		{ // 66
			String:                   ":\u002C\r\n",
			ExpectedString:            "\u002C",
			ExpectedWritten: int64(len("\u002C")),
		},
		{ // 67
			String:                   ":\u002D\r\n",
			ExpectedString:            "\u002D",
			ExpectedWritten: int64(len("\u002D")),
		},
		{ // 68
			String:                   ":\u002E\r\n",
			ExpectedString:            "\u002E",
			ExpectedWritten: int64(len("\u002E")),
		},
		{ // 69
			String:                   ":\u002F\r\n",
			ExpectedString:            "\u002F",
			ExpectedWritten: int64(len("\u002F")),
		},
		{ // 70
			String:                   ":\u0030\r\n",
			ExpectedString:            "\u0030",
			ExpectedWritten: int64(len("\u0030")),
		},
		{ // 71
			String:                   ":\u0031\r\n",
			ExpectedString:            "\u0031",
			ExpectedWritten: int64(len("\u0031")),
		},
		{ // 72
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









		{
			String:                   ":  \n",
			ExpectedString           : " ",
			ExpectedWritten: int64(len(" ")),
		},
		{
			String:                   ":  \r",
			ExpectedString           : " ",
			ExpectedWritten: int64(len(" ")),
		},
		{
			String:                   ":  \r\n",
			ExpectedString           : " ",
			ExpectedWritten: int64(len(" ")),
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
