package main

import (
	"math"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	l := len(src)
	if l > 2 {
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func(src2 []int64) {
			MergeSort(src2)
			wg.Done()
		}(src[:l/2])

		go func(src2 []int64) {
			MergeSort(src2)
			wg.Done()
		}(src[l/2:])

		wg.Wait()
	}

	if l > 2 {
		left := make([]int64, l/2 + 1)
		copy(left, src[:l/2])
		left[len(left)-1] = math.MaxInt64

		right := make([]int64, l - l/2 + 1)
		copy(right, src[l/2:])
		right[len(right)-1] = math.MaxInt64

		idxLeft, idxRight := 0, 0
		for i := range src {
			if left[idxLeft] < right[idxRight] {
				src[i] = left[idxLeft]
				idxLeft++
			}else {
				src[i] = right[idxRight]
				idxRight++
			}
		}
	} else if l == 2 {
		if src[0] > src[1] {
			src[0], src[1] = src[1], src[0]
		}
	}
}
