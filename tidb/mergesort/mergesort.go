package main

import "sync"

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	tmp := make([]int64, len(src))
	copy(tmp, src)
	merge(tmp, src, 0, len(src))
}

var sem = make(chan struct{}, 100)

func merge(src []int64, dst []int64, startIdx int, endIdx int) {
	l := endIdx - startIdx
	if l > 2 {
		half := startIdx + l/2

		var wg sync.WaitGroup

		select {
		case sem <- struct{}{}:
			{
				wg.Add(1)
				go func() {
					defer wg.Done()
					merge(dst, src, startIdx, half)
					<-sem
				}()
			}
		default:
			merge(dst, src, startIdx, half)
		}

		select {
		case sem <- struct{}{}:
			{
				wg.Add(1)
				go func() {
					defer wg.Done()
					merge(dst, src, half, endIdx)
					<-sem
				}()
			}
		default:
			merge(dst, src, half, endIdx)
		}

		wg.Wait()
		//merge(dst, src, startIdx, half)
		//merge(dst, src, half, endIdx)

		idxLeft, idxRight := startIdx, half
		for i := startIdx; i<endIdx; i++ {
			if idxRight == endIdx {
				dst[i] = src[idxLeft]
				idxLeft++
			}else if idxLeft == half {
				dst[i] = src[idxRight]
				idxRight++
			}else if src[idxLeft] < src[idxRight]{
				dst[i] = src[idxLeft]
				idxLeft++
			}else {
				dst[i] = src[idxRight]
				idxRight++
			}
		}
	}else if l == 2 {
		if src[startIdx] > src[startIdx+1] {
			dst[startIdx], dst[startIdx+1] = src[startIdx+1], src[startIdx]
		}
	}
}