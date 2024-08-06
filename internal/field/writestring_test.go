package field

import (
	"testing"

	"strings"
)

func TestWriteValue(t *testing.T) {

	tests := []struct{
		Name string
		Value string
		Expected string
	}{
		{
			Name:      "",
			Value:      "",
			Expected: ": \n",
		},



		{
			Name:     " ",
			Value:       "",
			Expected: " : \n",
		},
		{
			Name:     "  ",
			Value:        "",
			Expected: "  : \n",
		},
		{
			Name:     "   ",
			Value:         "",
			Expected: "   : \n",
		},
		{
			Name:     "    ",
			Value:          "",
			Expected: "    : \n",
		},



		{
			Name:     "",
			Value:      " ",
			Expected: ":  \n",
		},
		{
			Name:     "",
			Value:      "  ",
			Expected: ":   \n",
		},
		{
			Name:     "",
			Value:      "   ",
			Expected: ":    \n",
		},
		{
			Name:     "",
			Value:      "   ",
			Expected: ":    \n",
		},



		{
			Name:     "  ",
			Value:        "   ",
			Expected: "  :    \n",
		},



		{
			Name:     "name",
			Value:          "value",
			Expected: "name: value\n",
		},



		{
			Name:     "data",
			Value:          "once\ntwice\rthrice\r\nfource",
			Expected: "data: once\n"+
			          "data: twice\n"+
			          "data: thrice\n"+
			          "data: fource\n",
		},
	}

	for testNumber, test := range tests {

		var actualBuilder strings.Builder

		_, err := WriteString(&actualBuilder, test.Name, test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NAME: %q", test.Name)
			t.Logf("VALUE: %q", test.Value)
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
				t.Logf("NAME: %q", test.Name)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
