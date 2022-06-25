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
// - len   -append   -cap	-make
// create an empty slice with pre-specified length or capacity using ```make```

func main() {
	// test arrays
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(multi)

	var a []int
	a = append(a, 10, 15)
	fmt.Println(a)

	// empty slice with capacity of 5
	predef := make([]int, 5)
	predef[0] = 1
	predef[1] = 2
	predef[2] = 3
	predef[3] = 4
	predef[4] = 5
	predef = append(predef, 6, 7)
	fmt.Println(predef)
}