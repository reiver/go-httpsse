package anychar

import (
	"io"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-httpsse/internal/errors"
)


// Reads returns the next "any-char" character as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
//	any-char      = %x0000-0009 / %x000B-000C / %x000E-10FFFF
//	                ; a scalar value other than U+000A LINE FEED (LF) or U+000D CARRIAGE RETURN (CR)
//
// If the next character is NOT an "any-char", then it will return the err ErrNotAnyChar.
// (The next character will be unread.)
func Read(runescanner io.RuneScanner) (r rune, size int, err error) {
	if nil == runescanner {
		return 0, 0, errors.ErrNilRuneScanner
	}

	r, size, err = runescanner.ReadRune()

	if 0 < size && !isValid(r) {
		e := runescanner.UnreadRune()
		if nil != e {
			return r, size, erorr.Errorf("httpsse: problem unreading-rune %q (%U) that might have been a \"name-char\" but was not: %w", r, r, e)
		}

		return 0, 0, ErrNotAnyChar
	}

	return r, size, err
}


