package main

// the real star of Go's design is its implicit interfaces. it's the only abstract type in Go.
// the methods defined by an interface are called the method set of the interface.
// interfaces are named with "er" suffix and are declared in any block.
// the implicit behavior makes it interesting, enabling both type-safety and decoupling.

type LogicProvider struct {
}

func (lp LogicProvider) Process(data string) string {
	// business logic
	return "a string"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}