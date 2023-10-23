package main

import "fmt"

type bubbleSort struct{}

// Bubble Sort repeatedly swaps adjacent elements if out of order until the list is sorted.
func (b bubbleSort) Sort(numbers []int, debug bool) []int {
	// start from right side to left side
	// pick two neighboring elements
	countOperate := 0
	for {
		countOperate = 0
		for i := len(numbers) - 1; i-1 >= 0; i-- {
			if numbers[i-1] > numbers[i] {
				// if the element on the left bigger than the right, swap
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				countOperate++
			}
		}
		// if there nothing to swap in one round, then finish
		if countOperate == 0 {
			break
		}
		if debug {
			fmt.Printf("In-time sorted numbers: %v\n", numbers)
		}
	}
	return numbers
}
