package main

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

	directAddition := myAddition.AddTo
	fmt.Println(directAddition(15))
}