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

func (intTree *IntTree) Contains(value int) bool {
	switch {
	case intTree = nil:
		return false
	case value < intTree.value:
		return intTre.left.Contains(value)
	case value > intTree.value:
		return intTree.right.Contains(value)
	default:
		return true
	}
}