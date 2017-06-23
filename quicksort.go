package main

import "fmt"

func quicksort(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	low := []int{}
	high := []int{}
	equal := []int{}

	pivot := values[0]

	for _, val := range values {
		if val > pivot {
			high = append(high, val)
		} else if val < pivot {
			low = append(low, val)
		} else {
			equal = append(equal, val)
		}
	}
	return append(append(quicksort(low), equal...), quicksort(high)...)
}

func main() {
	fmt.Println(quicksort([]int{4, 3, 5, 6, 3, 2, 1}))
}
