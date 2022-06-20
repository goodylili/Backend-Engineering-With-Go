package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Naming variables (same as constants)
//var x int
//var y, z int
//var (
//	a int
//	b string
//	c float64
//)

// structs
// used for reading and writing data formats, and for grouping associative data
type PersonalID struct {
	name   string
	number int
	age    int
	isMale bool
}

// anonymous structs are used in table-driven tests

type PersonID struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	person1 := PersonalID{
		"Uzor",
		0001,
		15,
		true,
	}
	fmt.Println("First candidate: ", person1)

	newPerson := PersonID{
		"MacBobby",
		"theghostmac@gmail.com",
	}
	marsh, err := json.Marshal(newPerson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marsh))

	anotherPerson := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  "Chibuzor",
		Email: "macbobbychibuzor@gmaill.com",
	}
	anotherMarsh, err := json.Marshal(anotherPerson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(anotherMarsh))
}
