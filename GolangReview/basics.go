package main

import "fmt"

// Naming variables (same as constants)
//var x int
//var y, z int
//var (
//	a int
//	b string
//	c float64
//)

type PersonalID struct {
	name   string
	number int
	age    int
	isMale bool
}

func main() {
	person1 := PersonalID{
		"Uzor",
		0001,
		15,
		true,
	}

	fmt.Println("First candidate: ", person1)
}
