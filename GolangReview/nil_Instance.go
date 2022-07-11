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
}
