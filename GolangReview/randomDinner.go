package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	myDinner := []string{
		"fufu and egusi",
		"chicken and chips",
		"yam and egg sauce",
		"noodles and turkey",
	}

	rand.Seed(time.Now().Unix())
	fmt.Printf("I had %s for dinner tonight!", myDinner[rand.Intn(len(myDinner))])
}