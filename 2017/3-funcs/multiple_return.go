package main

import "log"

func Split(path string) (dir, file string) {
	// stub implementation
	dir = "directory"
	file = "filename"
	return
}

func main() {
	a, b := Split("some path")
	log.Printf("%q %q", a, b)
}
