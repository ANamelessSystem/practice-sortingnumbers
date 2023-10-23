package main

import "fmt"

type mergeSort struct{}

func (m mergeSort) Sort(numbers []int, debug bool) []int {
	// I struggled with recursion so hard, whatever this part finished by ChatGPT
	if len(numbers) <= 1 {
		return numbers
	}

	left := numbers[:len(numbers)/2]
	right := numbers[len(numbers)/2:]

	sortedLeft := m.Sort(left, debug)
	sortedRight := m.Sort(right, debug)

	// pretend the left and right are sorted (sub)arrays

	return mergeSortedArrays(sortedLeft, sortedRight, debug)
}

func mergeSortedArrays(left []int, right []int, debug bool) []int {
	// this function will merge two sorted arrays to one sorted array
	s := make([]int, len(left)+len(right))
	il, ir, is := 0, 0, 0

	// 3 situations: both, left remain, right remain
	// both arrays has unmerged elements
	for il < len(left) && ir < len(right) {
		if left[il] < right[ir] {
			s[is] = left[il]
			il++
			if debug {
				fmt.Printf("pick left: %v\n", s)
			}

		} else {
			s[is] = right[ir]
			ir++
			if debug {
				fmt.Printf("pick right: %v\n", s)
			}
		}
		is++
	}

	// after the above and left still has elements
	for il < len(left) {
		s[is] = left[il]
		il++
		is++
		if debug {
			fmt.Printf("pick remaining left: %v\n", s)
		}
	}

	// on the other hands
	for ir < len(right) {
		s[is] = right[ir]
		ir++
		is++
		if debug {
			fmt.Printf("pick remaining right: %v\n", s)
		}
	}
	return s
}
