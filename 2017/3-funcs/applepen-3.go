package main

type Pen struct{}

func (Pen) Write() {}

type Pineapple struct{}

func (Pineapple) Eat() {}

type PineapplePen struct {
	Pen
	Pineapple
}
