package endofline

import (
	"io"

	"github.com/reiver/go-ascii"
	"github.com/reiver/go-erorr"
	"github.com/reiver/go-utf8"

	"github.com/reiver/go-httpsse/internal/errors"
)

// Copy copies the 'end-of-line' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
//	end-of-line   = ( cr lf / cr / lf )
func Copy(writer io.Writer, runescanner io.RuneScanner) (written int64, err error) {
	if nil == runescanner {
		return written, errors.ErrNilRuneScanner
	}
	if nil == writer {
		return written, errors.ErrNilWriter
	}

	{
		var r0 rune
		{
			var size0 int

			r0, size0, err = runescanner.ReadRune()

			if size0 <= 0 {
				switch err {
				case nil:
					return written, erorr.Errorf("httpsse: problem reading rune (when trying to read first rune of end-of-line) — size = %d", size0)
				default:
					return written, erorr.Errorf("httpsse: problem reading rune (when trying to read first rune of end-of-line): %w", err)
				}
			}
		}

		if (ascii.CR != r0) && (ascii.LF != r0) {
			e := runescanner.UnreadRune()
			if nil != e {
				return written, erorr.Errorf("httpsse: problem unreading rune (when trying to read first rune of end-of-line): %w", e)
			}

			return written, ErrNotEndOfLine
		}

		{
			n, e := utf8.WriteRune(writer, r0)
			written += int64(n)
			if nil != e {
				return written, erorr.Errorf("httpsse: problem writing rune %q (%U) (when trying to read first rune of end-of-line): %w", r0, r0, e)
			}
		}

		if nil != err {
			return written, err
		}
	}

	{
		var r1 rune
		{
			var size1 int

			r1, size1, err = runescanner.ReadRune()

			if size1 <= 0 {
				switch err {
				case nil:
					return written, erorr.Errorf("httpsse: problem reading rune (when trying to second rune of end-of-line) — size = %d", size1)
				case io.EOF:
					return written, nil
				default:
					return written, erorr.Errorf("httpsse: problem reading rune (when trying to second rune of end-of-line): %w", err)
				}
			}
		}

		if ascii.LF != r1 {
			e := runescanner.UnreadRune()
			if nil != e {
				return written, erorr.Errorf("httpsse: problem unreading rune (when trying to read end-of-line): %w", e)
			}

			return written, nil
		}

		{
			n, e := utf8.WriteRune(writer, r1)
			written += int64(n)
			if nil != e {
				return written, erorr.Errorf("httpsse: problem writing rune %q (%U) (when trying to read first rune of end-of-line): %w", r1, r1, e)
			}
		}
	}

	return written, nil
}
