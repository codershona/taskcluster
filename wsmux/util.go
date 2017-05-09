package wsmux

import (
	"fmt"
	"io"
)

var (
	errMalformedHeader = fmt.Errorf("malformed header")
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getHeader(reader io.Reader) (header, error) {
	hdr := make([]byte, 5)
	l, err := reader.Read(hdr)

	if l != 5 {
		return nil, errMalformedHeader
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return hdr, nil
}

type Logger interface {
	Printf(format string, a ...interface{})
	Print(a ...interface{})
}

type nilLogger struct{}

func (n *nilLogger) Printf(format string, a ...interface{}) {}
func (n *nilLogger) Print(a ...interface{})                 {}
