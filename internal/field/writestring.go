package field

import (
	"io"
	"strings"

	 "github.com/reiver/go-httpsse/internal/errors"
)

// WriteString writes a 'field' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
// For example, this:
//
//	comment.WriteString(writer, "message", "Hello world!")
//
// Would write this:
//
//	"message: Hello world!\n"
//
// Also, for example, this:
//
//	comment.WriteString(writer, "do-it", "once\ntwice\nthrice\nfource")
//
// Would write this:
//
//	"do-it: once\ndo-it: twice\ndo-it: thrice\ndo-it: fource\n"
//
// I.e., this:
//
//	"do-it: once"   +"\n"+
//	"do-it: twice"  +"\n"+
//	"do-it: thrice" +"\n"+
//	"do-it: fource" +"\n"
func WriteString(writer io.Writer, name string, value string) (int, error) {
	if nil == writer {
		return 0, errors.ErrNilWriter
	}

	name = strings.ReplaceAll(name, "\r\n", " ")
	name = strings.ReplaceAll(name, "\r",   " ")
	name = strings.ReplaceAll(name, "\n",   " ")
	name = strings.ReplaceAll(name, ":",    "=")

	var prefix string = name +": "

	value = prefix + value
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\r", "\n")
	value = strings.ReplaceAll(value, "\n", "\n"+prefix)
	value += "\n"

	return io.WriteString(writer, value)
}
