package main

import "fmt"

// write a function to return the sum of all elements of a list of integers, with the list as an input argument
func sumInts(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total

}

func main() {
	var data = []int{20, 40, 60, 80, 100}
	fmt.Println(sumInts(data))
}
