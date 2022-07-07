package main

import (
	"fmt"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	if x == 0 && y == 0 {
		return 0
	} else {
		return x / y
	}
}

type operatorFuncType func(int, int) int

var operatorMap = map[string]operatorFuncType{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func main() {
	mathExpressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"6", "/", "3"},
		[]string{"5"},
		[]string{"two", "+", "three"},
	}

	for _, expressions := range mathExpressions {
		if len(expressions) != 3 {
			fmt.Println("invalid expression: ", expressions)
			continue
		}
		p1, err := strconv.Atoi(expressions[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expressions[1]
		operatorFunc, ok := operatorMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}
		p2, err := strconv.Atoi(expressions[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := operatorFunc(p1, p2)
		fmt.Println(result)
	}
}
