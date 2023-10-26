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

// This is an 'old' version of selection sort.
// In my previous implementation I use a int variable to storage the least number, which turns out that not necessary
// Also it has another redundant operation when linear search is finished and swapping the element, it only need to swap when the
// least number is not in the correct position.

// for p := 0; p < len(numbers); p++ {
// 	leastNumber := numbers[p]
// 	leastNumberIndex := p
// 	if debug {
// 		fmt.Printf("Picked number: %v\nNow position is: %v\n", leastNumber, p)
// 	}
// 	for i := p; i < len(numbers); i++ {
// 		if leastNumber > numbers[i] {
// 			leastNumber = numbers[i]
// 			leastNumberIndex = i
// 		}
// 	}
// 	// done with linear searching
// 	// the least number found, now swap the position with present element
// 	if debug {
// 		fmt.Printf("%d iteration, least number: %v\nIt's position is: %v\n", p, leastNumber, leastNumberIndex)
// 	}
// 	numbers[leastNumberIndex] = numbers[p]
// 	numbers[p] = leastNumber
// 	if debug {
// 		fmt.Printf("In-time sorted numbers: %v\n", numbers)
// 	}
// }

// This is a 'improved' version, which is pretty funny that it costs more time(5-10s) in random tests.
// First is the variable int 'leastNumber'
// it might be a duplication var but turns out accessing an array of 1000 elements is much heavier than frequently assigning values to a local variable.
// In general, local variables may be stored in CPU registers or caches, whereas array accesses may involve more memory accesses.
// The second is the swapping operation, in an array containing 1000 random elements,
// the probability of the minimum value being in exactly the right place is so low
// that judging this to avoid the overhead of swapping gains very little
// Improved: 50s Old: 40s

// for p := 0; p < len(numbers); p++ {
// 	leastIndex := p
// 	if debug {
// 		fmt.Printf("Picked number: %v\nNow position is: %v\n", numbers[leastIndex], p)
// 	}
// 	for i := p + 1; i < len(numbers); i++ {
// 		if numbers[leastIndex] > numbers[i] {
// 			// only need the least number's index
// 			leastIndex = i
// 		}
// 	}
// 	// done with linear searching
// 	// if the least number is not in its correct position, swap it with the current element
// 	// if leastIndex != p {
// 	// 	if debug {
// 	// 		fmt.Printf("%d iteration, least number: %v\nIt's position is: %v\n", p, numbers[leastIndex], leastIndex)
// 	// 	}
// 	// 	numbers[p], numbers[leastIndex] = numbers[leastIndex], numbers[p]
// 	// }

// 	// Theoretically it is better to check if you have to swap before you swap, but in practice it is more efficient to swap directly
// 	numbers[p], numbers[leastIndex] = numbers[leastIndex], numbers[p]
// 	if debug {
// 		fmt.Printf("In-time sorted numbers: %v\n", numbers)
// 	}
// }
