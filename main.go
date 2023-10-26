package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Sorter interface {
	Sort([]int, bool) []int
}

func main() {
	var algorithms = []string{"bubble", "quick", "merge", "insertion", "selection", "heap", "mysort"}

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Println("Use -a for algorithms and -r for additional random test.")
		fmt.Println("Available algorithms:", algorithms)
		flag.PrintDefaults()
	}

	var sorter Sorter

	algoPtr := flag.String("a", "", "sorting algorithm to use (e.g., bubble, quick, merge)")
	randomPtr := flag.Bool("r", false, "run random test with multiple arrays that each contains 1000 elements ")
	parallelPtr := flag.Bool("p", false, "run tests in parallel using goroutines")

	flag.Parse()

	validAlgo := false
	for _, algo := range algorithms {
		if *algoPtr == algo {
			validAlgo = true
			break
		}
	}

	if !validAlgo {
		fmt.Println("Unknown algorithm:", *algoPtr)
		fmt.Println("Available algorithms are:", algorithms)
		fmt.Println("Use -a for algorithms and -r for random test.")
		return
	}

	switch *algoPtr {
	case "bubble":
		sorter = bubbleSort{}
	case "quick":
		sorter = quickSort{}
	case "insertion":
		sorter = insertionSort{}
	case "merge":
		sorter = mergeSort{}
	case "selection":
		sorter = selectionSort{}
	case "heap":
		sorter = heapSort{}
	case "mysort":
		sorter = mySort{}
	default:
		fmt.Println("Unknown algorithm")
		return
	}

	if *randomPtr {
		if *parallelPtr {
			randomTestParallel(sorter, false)
		} else {
			randomTestSequential(sorter, false)
		}
	} else {
		simpleTest(sorter, true)
	}
}

func randomTestSequential(sorter Sorter, debug bool) {
	timeStart := time.Now()
	testCases := 1000
	for i := 0; i < testCases; i++ {
		numbers := generateRandomSlice(10000)
		sortedNumbers := sorter.Sort(numbers, debug)
		if !sort.IntsAreSorted(sortedNumbers) {
			fmt.Printf("Failed in the %d run\n", i)
			fmt.Printf("Failed detail:\nOrigin numbers: %v\nSorted numbers: %v\n", numbers, sortedNumbers)
			return
		}
	}
	fmt.Printf("All %d tests passed. Time elapsed: %v\n", testCases, time.Since(timeStart))
}

func randomTestParallel(sorter Sorter, debug bool) {
	testCases := 1000
	results := make(chan bool, testCases)
	timeStart := time.Now()

	for i := 0; i < testCases; i++ {
		go func() {
			numbers := generateRandomSlice(10000)
			sortedNumbers := sorter.Sort(numbers, debug)
			results <- sort.IntsAreSorted(sortedNumbers)
		}()
	}

	failed := false
	for i := 0; i < testCases; i++ {
		if !<-results {
			fmt.Printf("Failed in the %d run\n", i)
			failed = true
			break
		}
	}

	if !failed {
		fmt.Printf("All %d tests passed. Time elapsed: %v\n", testCases, time.Since(timeStart))
	}
}

func simpleTest(sorter Sorter, debug bool) {
	numbers := []int{1, 3, 9, 5, 4, 2, 6, 8, 7, 7}
	timeStart := time.Now()
	fmt.Printf("Origin numbers: %v\n", numbers)
	fmt.Printf("Sorted numbers: %v\n", sorter.Sort(numbers, debug))
	fmt.Printf("Time elapsed: %v\n", time.Since(timeStart))
}

func generateRandomSlice(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = r.Intn(1000)
	}
	return result
}
