package sortAlgo

import "sort"

func LibSort(arr []int64) []int64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
