package main

import "io"

// Discarder implements io.Writer.
// It can be used when you need to discard output.
type Discarder struct{}

func (Discarder) Write(b []byte) (int, error) {
	return len(b), nil
}

func main() {
	var w io.Writer
	w = Discarder{}
	w.Write([]byte("this will be discarded"))
}
