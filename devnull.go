package devnull

import (
	"io"
	"io/ioutil"
)

type reader struct{}

func (reader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

var Reader io.Reader

func Read(b []byte) (n int, err error) {
	return Reader.Read(b)
}

var Writer io.Writer

func Write(p []byte) (n int, err error) {
	return Writer.Write(p)
}

func init() {
	Reader = new(reader)
	Writer = ioutil.Discard
}
