package main

// a method value is like closure.

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(value int) int {
	return a.start + value
}

func main() {
	myAddition := Adder{
		start: 10,
	}
	fmt.Println(myAddition.AddTo(20))

	// method value
	directAddition := myAddition.AddTo
	fmt.Println(directAddition(15))

	// method expression
	anotherMethod := Adder.AddTo
	fmt.Println(anotherMethod(myAddition, 10))
}