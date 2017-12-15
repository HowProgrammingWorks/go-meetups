package main

import "fmt"

type Student struct {
	Name  string
	Group string
	Age   int
}

const n = 10

func makeSlice(students [n]Student, from, to int) ([]Student, error) {
	if from < 0 || from > to || to > len(students) {
		return nil, fmt.Errorf("invalid values of from and to")
	}
	res := make([]Student, 0, to-from)
	for i := from; i < to; i++ {
		res = append(res, students[i])
	}
	return res, nil
}
