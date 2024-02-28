package sortAlgo

import "sync"

// Goroutine for sorting an array
func SortingMulti(arr []int64, wg *sync.WaitGroup) {
	MergeSortMulti(arr)
	wg.Done()
}

func MergeSortMulti(arr []int64) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	leftArray := make([]int64, mid)
	rightArray := make([]int64, len(arr)-mid)

	copy(leftArray, arr[:mid])
	copy(rightArray, arr[mid:])

	MergeSortMulti(leftArray)
	MergeSortMulti(rightArray)

	Merge(arr, leftArray, rightArray)
}

func mergeSequency(left, right []int64) []int64 {
	ans := make([]int64, len(left)+len(right))
	Merge(ans, left, right)
	return ans
}

func MergeMultiArray(arr [][]int64) []int64 {
	ans := []int64{}
	for i := 0; i < len(arr); i++ {
		ans = mergeSequency(ans, arr[i])
	}
	return ans
}
