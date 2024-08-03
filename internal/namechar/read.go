package namechar

import (
	"io"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-httpsse/internal/errors"
)


// Reads returns the next "name-char" character as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
//	name-char     = %x0000-0009 / %x000B-000C / %x000E-0039 / %x003B-10FFFF
//	                ; a scalar value other than U+000A LINE FEED (LF), U+000D CARRIAGE RETURN (CR), or U+003A
//
// If the next character is NOT a "name-char", then it will return the err ErrNotNameChar.
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

		return 0, 0, ErrNotNameChar
	}

	return r, size, err
}


