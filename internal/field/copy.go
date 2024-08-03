package field

import (
	"io"

	"github.com/reiver/go-httpsse/internal/errors"
	"github.com/reiver/go-httpsse/internal/fieldvalue"
	"github.com/reiver/go-httpsse/internal/name"
)

// Copy copies the 'field' as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
//	field         = 1*name-char [ colon [ space ] *any-char ] end-of-line
func Copy(namewriter io.Writer, valuewriter io.Writer, runescanner io.RuneScanner) (namewritten int64, valuewritten int64, err error) {
	if nil == runescanner {
		return 0, 0, errors.ErrNilRuneScanner
	}
	if nil == namewriter {
		return 0, 0, errNilNameWriter
	}
	if nil == valuewriter {
		return 0, 0, errNilValueWriter
	}

	{
		var n int64

		n, err = name.Copy(namewriter, runescanner)
		namewritten += n

		switch err {
		case nil:
			// nothing here.
		case name.ErrNotName:
			return namewritten, valuewritten, ErrNotFieldBecauseNotName
		default:
			return namewritten, valuewritten, err
		}
	}

	{
		var n int64

		n, err = fieldvalue.Copy(valuewriter, runescanner)
		valuewritten += n

		if nil != err {
			return namewritten, valuewritten, err
		}
	}

	return namewritten, valuewritten, err
}
