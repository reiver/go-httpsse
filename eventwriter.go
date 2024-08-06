package httpsse

import (
	"io"

	"github.com/reiver/go-httpsse/internal/comment"
	"github.com/reiver/go-httpsse/internal/field"
)

type EventWriter interface {
	WriteComment(string) error
	WriteData(string) error
	WriteEvent(string) error
	WriteID(string) error
}

type internalEventWriter struct {
	writer io.Writer
}

var _ EventWriter = internalEventWriter{}

func (receiver internalEventWriter) WriteComment(value string) error {
	_, err := comment.WriteString(receiver.writer, value)
	return err
}

func (receiver internalEventWriter) WriteData(value string) error {
	_, err := field.WriteString(receiver.writer, "data", value)
	return err
}

func (receiver internalEventWriter) WriteEvent(value string) error {
	_, err := field.WriteString(receiver.writer, "event", value)
	return err
}

func (receiver internalEventWriter) WriteID(value string) error {
	_, err := field.WriteString(receiver.writer, "id", value)
	return err
}


