package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) personString() string {
	return fmt.Sprintf("%s %s of age %d", p.firstName, p.lastName, p.age)
}

func main() {

}