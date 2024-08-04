package httpsse

import (
	"bytes"
	"io"
	"strings"

	"github.com/reiver/go-erorr"

	"github.com/reiver/go-httpsse/internal/comment"
	"github.com/reiver/go-httpsse/internal/endofline"
	"github.com/reiver/go-httpsse/internal/field"
)

// readEvent reads a single event, and writes it to the EventWriter.
func readEvent(eventwriter EventWriter, runescanner io.RuneScanner) error {
	if nil == eventwriter {
		return errNilEventWriter
	}
	if nil == runescanner {
		return errNilRuneScanner
	}

	loop: for {
		{
			_, err := endofline.Copy(io.Discard, runescanner)

			switch err {
			case nil:
	/////////////////////// BREAK
				break loop
			case endofline.ErrNotEndOfLine:
				// nothing here.
			default:
				return erorr.Errorf("httpsse: problem trying to read potential end-of-line (EOL): %w", err)
			}
		}

		{
			_, err := comment.Copy(io.Discard, runescanner)

			switch err {
			case nil:
	/////////////////////// CONTINUE
				continue
			case comment.ErrNotComment:
				// nothing here.
			default:
				return erorr.Errorf("httpsse: problem trying to read potential comment: %w", err)
			}
		}

		{
			var nameBuffer strings.Builder
			var valueBuffer bytes.Buffer

			_, _, err := field.Copy(&nameBuffer, &valueBuffer, runescanner)

			switch err {
			case nil:
				// nothing here
			case field.ErrNotField:
				return erorr.Errorf("httpsse: problem trying to read potential field: %w", err)
			default:
				return erorr.Errorf("httpsse: problem trying to read potential field: %w", err)
			}

			var name string = nameBuffer.String()

			switch name {
			case "event":
				eventwriter.SetEventName(valueBuffer.String())
			case "id":
				eventwriter.SetEventID(valueBuffer.String())
			case "data":
				io.WriteString(eventwriter, valueBuffer.String())
				io.WriteString(eventwriter, "\n")
			case "retry":
				
//@TODO
				
			default:
				
//@TODO
				
			}
		}
	}

	return nil
}
