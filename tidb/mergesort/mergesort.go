package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	tmp := make([]int64, len(src))
	copy(tmp, src)
	merge(tmp, src, 0, len(src))
}

func merge(src []int64, dst []int64, startIdx int, endIdx int) {
	l := endIdx - startIdx
	if l > 2 {
		half := startIdx + l/2
		//var wg sync.WaitGroup
		//wg.Add(2)

		//go func(src1 []int64, dst1 []int64, startIdx1 int, endIdx1 int) {
		//	defer wg.Done()
		//	merge(src1, dst1, startIdx1, endIdx1)
		//}(dst, src, startIdx, half)
		//
		//go func(src2 []int64, dst2 []int64, startIdx2 int, endIdx2 int) {
		//	defer wg.Done()
		//	merge(src2, dst2, startIdx2, endIdx2)
		//}(dst, src, half, endIdx)

		//wg.Wait()
		merge(dst, src, startIdx, half)
		merge(dst, src, half, endIdx)

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