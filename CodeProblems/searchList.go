package main

// write a function that will search a list for a particular value

func SearchList(data []int, value int) bool {
	size := len(data)
	for index := 0; index < size; index++ {
		if value == data[index] {
			return true
		}
	}
	return false
}
