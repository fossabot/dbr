package dbr

import "bytes"

// Buffer collects strings, and values that are ready to be interpolated.
type Buffer interface {
	WriteString(string) (int, error)
	String() string

	WriteValue(v ...interface{}) (err error)
	Value() []interface{}
}

type buffer struct {
	bytes.Buffer
	v []interface{}
}

// NewBuffer creates a new Buffer.
func NewBuffer() Buffer {
	return &buffer{}
}

func (b *buffer) WriteValue(v ...interface{}) error {
	b.v = append(b.v, v...)
	return nil
}

func (b *buffer) Value() []interface{} {
	return b.v
}
