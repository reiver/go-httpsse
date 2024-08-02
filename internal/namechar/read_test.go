package namechar

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-utf8"
)

func TestRead(t *testing.T) {

	tests := []struct{
		Rune rune
		ExpectedRune rune
		ExpectedSize int
	}{
		{
			Rune:         '\u0000',
			ExpectedRune: '\u0000',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0001',
			ExpectedRune: '\u0001',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0002',
			ExpectedRune: '\u0002',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0003',
			ExpectedRune: '\u0003',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0004',
			ExpectedRune: '\u0004',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0005',
			ExpectedRune: '\u0005',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0006',
			ExpectedRune: '\u0006',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0007',
			ExpectedRune: '\u0007',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0008',
			ExpectedRune: '\u0008',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0009',
			ExpectedRune: '\u0009',
			ExpectedSize: 1,
		},



		{
			Rune:         '\u000B',
			ExpectedRune: '\u000B',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u000C',
			ExpectedRune: '\u000C',
			ExpectedSize: 1,
		},



		{
			Rune:         '\u000E',
			ExpectedRune: '\u000E',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u000F',
			ExpectedRune: '\u000F',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0010',
			ExpectedRune: '\u0010',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0011',
			ExpectedRune: '\u0011',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0012',
			ExpectedRune: '\u0012',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0013',
			ExpectedRune: '\u0013',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0014',
			ExpectedRune: '\u0014',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0015',
			ExpectedRune: '\u0015',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0016',
			ExpectedRune: '\u0016',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0017',
			ExpectedRune: '\u0017',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0018',
			ExpectedRune: '\u0018',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0019',
			ExpectedRune: '\u0019',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u001A',
			ExpectedRune: '\u001A',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u001B',
			ExpectedRune: '\u001B',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u001C',
			ExpectedRune: '\u001C',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u001D',
			ExpectedRune: '\u001D',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u001E',
			ExpectedRune: '\u001E',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u001F',
			ExpectedRune: '\u001F',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0020',
			ExpectedRune: '\u0020',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0021',
			ExpectedRune: '\u0021',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0022',
			ExpectedRune: '\u0022',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0023',
			ExpectedRune: '\u0023',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0024',
			ExpectedRune: '\u0024',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0025',
			ExpectedRune: '\u0025',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0026',
			ExpectedRune: '\u0026',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0027',
			ExpectedRune: '\u0027',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0028',
			ExpectedRune: '\u0028',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0029',
			ExpectedRune: '\u0029',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u002A',
			ExpectedRune: '\u002A',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u002B',
			ExpectedRune: '\u002B',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u002C',
			ExpectedRune: '\u002C',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u002D',
			ExpectedRune: '\u002D',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u002E',
			ExpectedRune: '\u002E',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u002F',
			ExpectedRune: '\u002F',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0030',
			ExpectedRune: '\u0030',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0031',
			ExpectedRune: '\u0031',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0032',
			ExpectedRune: '\u0032',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0033',
			ExpectedRune: '\u0033',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0034',
			ExpectedRune: '\u0034',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0035',
			ExpectedRune: '\u0035',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0036',
			ExpectedRune: '\u0036',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0037',
			ExpectedRune: '\u0037',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0038',
			ExpectedRune: '\u0038',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0039',
			ExpectedRune: '\u0039',
			ExpectedSize: 1,
		},



		{
			Rune:         '\u003B',
			ExpectedRune: '\u003B',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u003C',
			ExpectedRune: '\u003C',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u003D',
			ExpectedRune: '\u003D',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u003E',
			ExpectedRune: '\u003E',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u003F',
			ExpectedRune: '\u003F',
			ExpectedSize: 1,
		},
		{
			Rune:         '\u0040',
			ExpectedRune: '\u0040',
			ExpectedSize: 1,
		},



		// ...



		{
			Rune:         '\U0010FFFD',
			ExpectedRune: '\U0010FFFD',
			ExpectedSize: 4,
		},
		{
			Rune:         '\U0010FFFE',
			ExpectedRune: '\U0010FFFE',
			ExpectedSize: 4,
		},
		{
			Rune:         '\U0010FFFF',
			ExpectedRune: '\U0010FFFF',
			ExpectedSize: 4,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(string(test.Rune))
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualRune, actualSize, err := Read(runescanner)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.ExpectedRune
			actual   :=        actualRune

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q (%U)", expected, expected)
				t.Logf("ACTUAL:   %q (%U)", actual, actual)
				continue
			}
		}

		{
			expected := test.ExpectedSize
			actual   :=        actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}
	}
}

func TestRead_errNotNameChar(t *testing.T) {

	tests := []struct{
		Rune rune
	}{
		{
			Rune: '\u000A',
		},



		{
			Rune: '\u000D',
		},



		{
			Rune: '\u003A',
		},
	}

	for testNumber, test := range tests {
		var reader io.Reader = strings.NewReader(string(test.Rune))
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		actualRune, actualSize, err := Read(runescanner)

		{
			var expected error = ErrNotNameChar
			actual := err

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%T) %s", expected, expected)
				t.Logf("ACTUAL:   (%T) %s", actual, actual)
				continue
			}
		}

		{
			var expected rune = 0
			actual   := actualRune

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q (%U)", expected, expected)
				t.Logf("ACTUAL:   %q (%U)", actual, actual)
				continue
			}
		}

		{
			expected := 0
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}
	}
}
