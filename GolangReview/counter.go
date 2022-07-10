package main

import "time"

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c Counter) Increment() {
	c.total++
	c.lastUpdated
}
