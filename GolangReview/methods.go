package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) customerData() string {
	return fmt.Sprintf("Customer name: %s %s. Customer age: %d", p.firstName, p.lastName, p.age)
}

func main() {
	customer1 := Person{
		"Bryan",
		"Goldings",
		47,
	}

	feedData := customer1.customerData()
	fmt.Println(feedData)
}