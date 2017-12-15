package main

import (
	"log"
	"time"
)

func main() {
	arr := []int{10, 12, 15, 21}
	funcs := make([]func(), len(arr))
	for i := range arr {
		funcs[i] = func() {
			log.Printf("Index: %d, element: %v", i, arr[i])
		}
	}

	time.Sleep(3 * time.Second)
	for _, f := range funcs {
		f()
	}
}
