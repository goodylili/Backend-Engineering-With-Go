package main

// write a function to return the sum of all elements of a list of integers, with the list as an input argument

func SumInts(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total

}
