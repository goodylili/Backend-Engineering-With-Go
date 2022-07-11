package main

import "fmt"

type IntTree struct {
	value int
	left  *IntTree
	right *IntTree
}

func (intTree *IntTree) Insert(value int) *IntTree {
	if intTree == nil {
		return &IntTree{value: value}
	}
	if value < intTree.value {
		intTree.left = intTree.left.Insert(value)
	} else if value > intTree.value {
		intTree.right = intTree.right.Insert(value)
	}
	return intTree
}

func (intTree *IntTree) Contains(value int) bool {
	switch {
	case intTree == nil:
		return false
	case value < intTree.value:
		return intTree.left.Contains(value)
	case value > intTree.value:
		return intTree.right.Contains(value)
	default:
		return true
	}
}

func main() {
	var intTree *IntTree
	intTree = intTree.Insert(5)
	intTree = intTree.Insert(3)
	intTree = intTree.Insert(10)
	intTree = intTree.Insert(2)
	intTree = intTree.Insert(4)
	fmt.Println(intTree.Contains(2))
	fmt.Println(intTree.Contains(12))
}
