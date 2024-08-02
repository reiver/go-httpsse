package endofline

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"
)

func TestCopy(t *testing.T) {

	tests := []struct{
		String string
		Expected string
	}{
		{
			String:   "\n",
			Expected: "\n",
		},
		{
			String:   "\r",
			Expected: "\r",
		},
		{
			String:   "\r\n",
			Expected: "\r\n",
		},



		{
			String:   "\nsomething",
			Expected: "\n",
		},
		{
			String:   "\rsomething",
			Expected: "\r",
		},
		{
			String:   "\r\nsomething",
			Expected: "\r\n",
		},



		{
			String:   "\n\r",
			Expected: "\n",
		},
		{
			String:   "\r\r",
			Expected: "\r",
		},
		{
			String:   "\r\n\r",
			Expected: "\r\n",
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		var reader io.Reader = strings.NewReader(test.String)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		written, err := Copy(&actualBuffer, runescanner)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.Expected
			actual := actualBuffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual end-of-line rune(s) written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := int64(len(test.Expected))
			actual := written

			if expected != actual {
				t.Errorf("For test #%d, the actual number-of-bytes for the end-of-line rune(s) written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}

func TestCopy_errNotEndOfLine(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String:   "\u0000",
		},



		{
			String:   "a",
		},
	}

	for testNumber, test := range tests {

		var actualBuffer strings.Builder

		var reader io.Reader = strings.NewReader(test.String)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		written, err := Copy(&actualBuffer, runescanner)

		{
			expected := int64(0)
			actual := written

			if expected != actual {
				t.Errorf("For test #%d, the actual number-of-bytes for the end-of-line rune(s) written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			var expected error = ErrNotEndOfLine
			actual := err

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
