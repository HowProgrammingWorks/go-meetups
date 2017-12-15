package main

type Apple struct{}

func (Apple) FeedGodOfDeath() {}

type Pen struct{}

func (Pen) Write() {}

type ApplePen struct {
	Pen
	Apple
}
