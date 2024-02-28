package sortAlgo

// MergeSort algo
func MergeSort(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := make([]int64, mid)
	right := make([]int64, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	MergeSort(left)
	MergeSort(right)

	Merge(arr, left, right)
	return arr
}

func Merge(arr, left, right []int64) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
