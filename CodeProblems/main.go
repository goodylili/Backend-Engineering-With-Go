package main

import (
	"fmt"
	//"github.com/theghostmac/Backend-Engineering-With-Go"
)

func main() {
	// from sumlist.go
	var data = []int{20, 40, 60, 80, 100}
	fmt.Println(SumInts(data))
	// from searchList.go
	fmt.Println(SearchList(data, 40))
}
