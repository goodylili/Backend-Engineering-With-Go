package main

import "fmt"

// arrays
// are too rigid to use, so are not often used
var x = [3]int{1, 2, 3}

// assignment can be done by index in a function
// size can be omitted
var y = [...]int{10, 20, 30}

// multi dimensional arrays exist
var multi = [3][3]int{}

// slices
// same initialization as arrays, but without size
// operations include
// - len   -append   -cap

func main() {
	// test arrays
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(multi)

	var a []int
	a = append(a, 10, 15)
	fmt.Println(a)
}