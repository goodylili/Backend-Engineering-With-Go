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
// use make when you know how large you need the slice to be but don't know the values to populate with

// strings
// strings are sequences of utf-8 bytes
var s string = "Hello there"
var b byte = s[6]

// slicing and indexing also works for strings. only use it when your string contains characters using one byte
// most data in Go are read and written in bytes. type conversion of strings to byte slices are common.

// maps
// maps are used to associate one value with another.
var nilMap map[string]int

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

	receiverSlice := make([]int, 6)
	newSlice := copy(receiverSlice, predef)
	fmt.Println(receiverSlice, newSlice)

	fmt.Println(s, b)
	var byteOfStringS []byte = []byte(s)
	var runeOfStringS []rune = []rune(s)
	fmt.Println(byteOfStringS)
	fmt.Println(runeOfStringS)

	// maps
	fmt.Println(nilMap)
	totalWins := map[string]int{
		"Orcas":   10,
		"Lions":   12,
		"Kittens": 11,
	}
	fmt.Println(totalWins)
}
