package main

import "fmt"

type quickSort struct{}

func (q quickSort) Sort(numbers []int, debug bool) []int {
	if debug {
		fmt.Printf("=========================\nStart to handle %v\n", numbers)
	}
	if len(numbers) <= 1 {
		return numbers
	}
	pivot := len(numbers) - 1
	l := 0
	r := pivot - 1
	if debug {
		fmt.Printf("pivot: %v(%v)\nl: %v(%v)\nr: %v(%v)\n", pivot, numbers[pivot], l, numbers[l], r, numbers[r])
	}
	for l < pivot {
		if numbers[l] >= numbers[pivot] {
			if debug {
				fmt.Printf("l(hold): %v(%v)\n", l, numbers[l])
			}
			for l < r {
				if numbers[r] < numbers[pivot] {
					if debug {
						fmt.Printf("swap: %v(%v) <==> %v(%v)\n", l, numbers[l], r, numbers[r])
					}
					numbers[l], numbers[r] = numbers[r], numbers[l]
					if debug {
						fmt.Printf("In-time array: %v\n", numbers)
					}
					break
				} else {
					r--
					if debug {
						fmt.Printf("r move: %v(%v)\n", r, numbers[r])
					}
				}
			}
			if l == r {
				if debug {
					fmt.Printf("r met l(hold) at: %v(%v)\nswap: %v(%v) <==> %v(%v)\n", l, numbers[l], l, numbers[l], pivot, numbers[pivot])
				}
				numbers[l], numbers[pivot] = numbers[pivot], numbers[l]
				if debug {
					fmt.Printf("In-time array: %v\n", numbers)
				}
				break
			}
		} else {
			l++
			if debug {
				fmt.Printf("l move: %v(%v)\n", l, numbers[l])
			}
		}
	}
	if debug {
		fmt.Printf("End an iteration, the pivot %v(%v) has been sorted.\n", l, numbers[l])
	}

	leftPart := numbers[:l]
	q.Sort(leftPart, debug)
	if l+1 < len(numbers) {
		rightPart := numbers[l+1:]
		q.Sort(rightPart, debug)
	}
	return numbers
}
