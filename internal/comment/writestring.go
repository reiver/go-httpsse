package comment

import (
	"io"
	"strings"

	 "github.com/reiver/go-httpsse/internal/errors"
)

// WriteString writes a 'comment' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
// For example, this:
//
//	comment.WriteString(writer, "Hello world!")
//
// Would write this:
//
//	":Hello world!\n"
//
// Also, for example, this:
//
//	comment.WriteString(writer, "once\ntwice\nthrice\nfource")
//
// Would write this:
//
//	":once\n:twice\n:thrice\n:fource\n"
//
// I.e., this:
//
//	":once"   +"\n"+
//	":twice"  +"\n"+
//	":thrice" +"\n"+
//	":fource" +"\n"
func WriteString(writer io.Writer, str string) (int, error) {
	if nil == writer {
		return 0, errors.ErrNilWriter
	}

	var value string = ":" + str
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\r", "\n")
	value = strings.ReplaceAll(value, "\n", "\n:")
	value += "\n"

	return io.WriteString(writer, value)
}
