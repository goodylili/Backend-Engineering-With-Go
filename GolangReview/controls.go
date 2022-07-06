package main

import "fmt"

func main() {
	// FizzBuzz Algorithm
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Fizz")
			continue
		} else {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
