package main

import "fmt"

type mySort struct{}

func (m mySort) Sort(numbers []int, debug bool) []int {
	sortedNumbers := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		selectedNumber := numbers[i]
		if debug {
			fmt.Printf("Selected number: %v\n", selectedNumber)
		}
		seq := 0
		for _, v := range numbers {
			if selectedNumber > v {
				seq++
			}
		}
		if debug {
			fmt.Printf("Bigger than selected number: %v\n", seq)
		}
		for j := 0; j < len(numbers); j++ {
			if sortedNumbers[seq+j] == 0 {
				sortedNumbers[seq+j] = selectedNumber
				if debug {
					fmt.Printf("Offset of selected number: %v\n", j)
				}
				break
			}
		}
		if debug {
			fmt.Printf("In-time sorted numbers: %v\n", sortedNumbers)
		}
	}
	return sortedNumbers
}
