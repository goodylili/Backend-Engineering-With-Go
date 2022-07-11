package main

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
