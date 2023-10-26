package main

import "fmt"

type mergeSort struct{}

// Split int array into two arrays until sorted or has no elements
func (m mergeSort) Sort(numbers []int, debug bool) []int {
	// I struggled with recursion so hard, whatever this part finished by ChatGPT
	// if numbers only has one or none element, it can treat to be sorted
	if len(numbers) <= 1 {
		return numbers
	}

	// split into two parts
	left := numbers[:len(numbers)/2]
	right := numbers[len(numbers)/2:]

	// sort both parts
	sortedLeft := m.Sort(left, debug)
	sortedRight := m.Sort(right, debug)

	// the left and right are already sorted, now merge them
	return mergeSortedArrays(sortedLeft, sortedRight, debug)
}

// Merge two sorted arrays to one sorted array
func mergeSortedArrays(left []int, right []int, debug bool) []int {
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
