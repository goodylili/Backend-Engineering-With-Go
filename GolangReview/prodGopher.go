package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dinners := []string{
		"chicken and chips",
		"fufu and egusi",
		"yam and egg sauce",
		"noodles and turkey",
	}

	rand.Seed(time.Now().Unix())

	fmt.Printf("I had %s for dinner tonight!", dinners[rand.Intn(len(dinners))])
}