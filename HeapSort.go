package main

type heapSort struct{}

func (h heapSort) Sort(numbers []int, debug bool) []int {
	// init heap (big-top)
	l := len(numbers)
	for i := l/2 - 1; i >= 0; i-- {
		numbers = heapAdjust(numbers, i, l)
	}

	// pull out the heap top to the array's end, and assume it has been sorted
	for i := l - 1; i > 0; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]

		// the node 0 had been changed, as the same logic in heapAdjust(), only adjust this specific branch
		heapAdjust(numbers, 0, i)
	}

	// // loop over, all elements sorted, biggest last
	return numbers
}

// build a BIG top heap here, n is the range of heap
// node[i] should bigger than node[2i+1] and node[2i+2]
func heapAdjust(numbers []int, i int, n int) []int {
	// use n check left and right node exists
	topIndex := i
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2

	if leftIndex < n && numbers[leftIndex] > numbers[topIndex] {
		topIndex = leftIndex
	}

	if rightIndex < n && numbers[rightIndex] > numbers[topIndex] {
		topIndex = rightIndex
	}

	// only adjust when parent node index changed
	if topIndex != i {
		numbers[i], numbers[topIndex] = numbers[topIndex], numbers[i]
		heapAdjust(numbers, topIndex, n)
	}

	return numbers
}
