package main

import (
	"log"
	"time"
)

func main() {
	arr := []int{10, 12, 15, 21}
	for i := range arr {
		go func(i int) {
			time.Sleep(3 * time.Second)
			log.Printf("Index: %d, element: %v", i, arr[i])
		}(i)
	}
	time.Sleep(4 * time.Second)
}
