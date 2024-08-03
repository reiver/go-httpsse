package fieldvalue

import (
	"io"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-httpsse/internal/anychars"
	"github.com/reiver/go-httpsse/internal/errors"
	"github.com/reiver/go-httpsse/internal/endofline"
)

// Copy copies the 'field-value' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
// Note that 'field-value' is implicitly defined.
//
//	[ colon [ space ] *any-char ] end-of-line
func Copy(writer io.Writer, runescanner io.RuneScanner) (written int64, err error) {
	if nil == runescanner {
		return 0, errors.ErrNilRuneScanner
	}
	if nil == writer {
		return 0, errors.ErrNilWriter
	}

	var r0 rune
	{
		var size int

		r0, size, err = runescanner.ReadRune()

		if io.EOF == err {
			return 0, io.EOF
		}

		if nil != err {
			return 0, erorr.Errorf("httpsse: problem reading rune: %w", err)
		}

		if size <= 0 {
			return 0, erorr.Errorf("httpsse: problem reading rune — size = %d", size)
		}

		if ':' != r0 {
			e := runescanner.UnreadRune()
			if nil != e {
				
			}
		}
	}

	if ':' == r0 {
		var r1 rune
		var size int

		r1, size, err = runescanner.ReadRune()

		if io.EOF == err {
			return 0, io.EOF
		}

		if nil != err {
			return 0, erorr.Errorf("httpsse: problem reading rune: %w", err)
		}

		if size <= 0 {
			return 0, erorr.Errorf("httpsse: problem reading rune — size = %d", size)
		}

		if ' ' != r1 {
			e := runescanner.UnreadRune()
			if nil != e {
				
			}
		}
	}

	{
		written, err = anychars.Copy(writer, runescanner)

		switch err {
		case nil:
			// nothing here.
		default:
			return written, err
		}
	}

	{
		_, err = endofline.Copy(io.Discard, runescanner)

		switch err {
		case nil:
			// nothing here
		case endofline.ErrNotEndOfLine:
			return written, ErrNotFieldValue
		default:
			return written, err
		}
	}

	return written, err
}
