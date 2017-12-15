package main

import (
	"fmt"
	"io"
)

// START OMIT

// ZeroReader implements io.Reader.
// It reads zero values.
type ZeroReader struct{}

func (ZeroReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

func main() {
	var r io.Reader
	s := make([]byte, 42)
	r = ZeroReader{}

	r.Read(s)

	fmt.Println(s)
}

// END OMIT
