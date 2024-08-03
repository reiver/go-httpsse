package comment

import (
	"io"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-httpsse/internal/anychars"
	"github.com/reiver/go-httpsse/internal/errors"
	"github.com/reiver/go-httpsse/internal/endofline"
)

// Copy copies the 'comment' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
//	comment       = colon *any-char end-of-line
//
// If the line is not a comment, it return ErrNotComment
func Copy(writer io.Writer, runescanner io.RuneScanner) (written int64, err error) {
	if nil == runescanner {
		return 0, errors.ErrNilRuneScanner
	}
	if nil == writer {
		return 0, errors.ErrNilWriter
	}

	{
		r, size, err := runescanner.ReadRune()

		if io.EOF == err {
			return 0, io.EOF
		}

		if nil != err {
			return 0, erorr.Errorf("httpsse: problem reading rune: %w", err)
		}

		if size <= 0 {
			return 0, erorr.Errorf("httpsse: problem reading rune — size = %d", size)
		}

		if ':' != r {
			return 0, ErrNotComment
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
			return written, ErrNotComment
		default:
			return written, err
		}
	}

	return written, err
}
