package safebufio

import (
	"bufio"
	"io"
	"sync"
)


type SafeReadWriter struct {
	mutex sync.Mutex
	value *bufio.ReadWriter
}

func (receiver *SafeReadWriter) Set(value *bufio.ReadWriter) {

	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	receiver.value = value
}

func (receiver *SafeReadWriter) Write(fn func(writer io.Writer)error) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.value {
		return errNilBufferedReadWriter
	}

	err := fn(receiver.value)

	receiver.value.Flush()

	return err
}
