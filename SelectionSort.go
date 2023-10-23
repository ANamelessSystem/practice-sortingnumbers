package main

import "fmt"

type selectionSort struct{}

func (s selectionSort) Sort(numbers []int, debug bool) []int {
	for p := 0; p < len(numbers); p++ {
		leastNumber := numbers[p]
		leastNumberIndex := p
		if debug {
			fmt.Printf("Picked number: %v\nNow position is: %v\n", leastNumber, p)
		}
		for i := p; i < len(numbers); i++ {
			if leastNumber > numbers[i] {
				leastNumber = numbers[i]
				leastNumberIndex = i
			}
		}
		// done with linear searching
		// the least number found, now swap the position with present element
		if debug {
			fmt.Printf("%d iteration, least number: %v\nIt's position is: %v\n", p, leastNumber, leastNumberIndex)
		}
		numbers[leastNumberIndex] = numbers[p]
		numbers[p] = leastNumber
		if debug {
			fmt.Printf("In-time sorted numbers: %v\n", numbers)
		}
	}
	return numbers
}
