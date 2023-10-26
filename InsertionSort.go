package main

import "fmt"

type insertionSort struct{}

func (i insertionSort) Sort(numbers []int, debug bool) []int {
	for i := 1; i < len(numbers); i++ {
		key := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > key {
			// if a number bigger than key was found in sorted part, move their position by 1 until found a number smaller or equal key
			// this loop is for moving numbers back
			numbers[j+1] = numbers[j]
			j = j - 1
			if debug {
				fmt.Printf("In-time sorted numbers: %v\n", numbers)
			}
		}
		// after moving every bigger number back, use the temporary storage number to cover the threshold+1 value
		numbers[j+1] = key
	}
	return numbers
}

// In my previous implementation
// I compared the key(numbers[i]) with elements from its immediate left (numbers[i-1]) to the first element and swap it when it meet the condition
// It's not an approach of insertion sorting, more like bubble sorting.

// The main issue is that I swap every elements that match the conditions immediately.
// This approach resulted in unnecessary operations (comparisons and swaps).
// The insertion approach is take the key out, find the right place to insert it, then adjusts the rest elements position.
// Also, in a typical insertion sort, it needs a temporary space to help storing the key.
// For a 1000 arrays test it cost about 42.4s, after improved efficiency it cost 12.9s

// for i := 1; i < len(numbers); i++ {
// 	for j := i - 1; j >= 0; j-- {
// 		// number[i] is the operating element, number[j](s) are sorted numbers
// 		// i will bigger than j by 1 initially
// 		if numbers[j+1] < numbers[j] {
// 			numbers[j+1], numbers[j] = numbers[j], numbers[j+1]
// 		}
// 	}

// 	if debug {
// 		fmt.Printf("In-time sorted numbers: %v\n", numbers)
// 	}
// }
