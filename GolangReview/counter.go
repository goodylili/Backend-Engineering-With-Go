package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// pointer receiver method

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// value receiver method

func (c Counter) MakeString() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.MakeString())
	c.Increment()
	fmt.Println(c.MakeString())
}
