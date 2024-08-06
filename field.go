package httpsse

import (
	"strings"
)

type internalField struct {
	name string
	values []string
}

func (receiver *internalField) Append(str string) {
	if nil == receiver {
		return
	}

	receiver.values = append(receiver.values, str)
}

func (receiver internalField) Collapse() string {
	return strings.Join(receiver.values, "\n")
}

func (receiver internalField) Name() string {
	return receiver.name
}

func (receiver *internalField) Reset() {
	receiver.name = ""
	receiver.values = nil
}

func (receiver internalField) String() string {
	return receiver.Collapse()
}

func (receiver internalField) Values() []string {
	return append([]string(nil), receiver.values...)
}
