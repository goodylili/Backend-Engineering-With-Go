package main

import (
	"errors"
	"fmt"
	"os"
)

func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func divAndRem(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot be divided by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result := div(20, 10)
	fmt.Println(result)

	res, remainder, err := divAndRem(50, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res, remainder)
}
