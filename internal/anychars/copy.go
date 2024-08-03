package anychars

import (
	"io"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-utf8"

	"github.com/reiver/go-httpsse/internal/anychar"
)

// Copy copies the 'any-chars' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
// Note that 'any-chars' is implicitly defined.
//
//	*any-char
func Copy(writer io.Writer, runescanner io.RuneScanner) (written int64, err error) {
	if nil == runescanner {
		return 0, errNilRuneScanner
	}
	if nil == writer {
		return 0, errNilWriter
	}

	for {
		r, size, err := anychar.Read(runescanner)

		if 0 < size {
			n, e := utf8.WriteRune(writer, r)
			written += int64(n)
			if nil != e {
				return written, erorr.Errorf("httpsse: problem write rune %q (%U) from event 'any-chars': %w", r, r, e)
			}
		}

		if nil != err {
			switch err {
			case anychar.ErrNotAnyChar:
				return written, nil
			case io.EOF:
				return written, ErrUnexpectedEOF
			default:
				return written, erorr.Errorf("httpsse: problem reading rune for event field 'any-chars': %w", err)
			}
		}
	}
}
