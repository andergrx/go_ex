package main

import (
	"fmt"
	"slices"
)

func main() {

	values := []int{1, 4, 7, 9, 11, 20, 33}
	fmt.Println("vals:", values)

	fmt.Println("low, high:", 6, "-", getLowIndex(6, values), getHighIndex(6, values))
	fmt.Println("low, high:", 0, "-", getLowIndex(0, values), getHighIndex(0, values))
	fmt.Println("low, high:", 99, "-", getLowIndex(99, values), getHighIndex(99, values))
}

func getLowIndex(val int, input []int) int {
	index, found := slices.BinarySearch(input, val)
	if found || index <= 0 {
		return index
	} else {
		return index - 1
	}
}

func getHighIndex(val int, input []int) int {
	index, found := slices.BinarySearch(input, val)
	if found {
		if index+1 >= len(input) {
			return index
		}
		return index + 1

	} else {
		if index >= len(input) {
			return index - 1
		}
		return index
	}
}
