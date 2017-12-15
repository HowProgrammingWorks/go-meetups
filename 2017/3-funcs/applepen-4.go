package main

type ApplePen struct {
	Pen
	Apple
}

type PineapplePen struct {
	Pen
	Pineapple
}

type PenPineappleApplePen struct {
	ApplePen
	PineapplePen
}
