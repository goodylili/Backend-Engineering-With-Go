package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// pointer receiver method used because I am modifying (also necessary when handling nil instances)

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// value receiver method

func (c Counter) MakeString() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doingUpdateWrong: ", c.MakeString())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doingUpdateRight:", c.MakeString())
}

func main() {
	var c Counter
	fmt.Println(c.MakeString())
	c.Increment()
	fmt.Println(c.MakeString())

	doUpdateWrong(c)
	fmt.Println("in main of doUpdateWrong: ", c.MakeString())
	doUpdateRight(&c)
	fmt.Println("in main of doUpdateRight: ", c.MakeString())
}
