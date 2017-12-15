package main

import (
	"time"
)

func runAfter(fn func() error, timeout time.Duration) {
	go func() {
		time.Sleep(timeout)
		fn()
	}()
}

const semestr = 3 * time.Second

type Lector struct{}

func (Lector) Otchislit(students []Student) {
	for _, student := range students {
		runAfter(student.GoToArmy, semestr)
	}
}
