package main

import "fmt"

type insertionSort struct{}

func (i insertionSort) Sort(numbers []int, debug bool) []int {
	for i := 1; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			// number[i] is the operating element, number[j](s) are sorted numbers
			// i will bigger than j by 1 initially
			if numbers[j+1] < numbers[j] {
				numbers[j+1], numbers[j] = numbers[j], numbers[j+1]
			}
		}
		if debug {
			fmt.Printf("In-time sorted numbers: %v\n", numbers)
		}
	}
	return numbers
}
