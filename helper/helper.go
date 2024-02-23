package helper

/*
Input:
  - numeberOfArray	: number of array that separated		Type: int
  - arrayInput		: array that separated					Type: []int64

Output:

	2-D array 													Type: [][]int64
*/
func SeparateArray(numberOfArray int, arrayInput []int64) [][]int64 {
	arrayLength := len(arrayInput)

	if numberOfArray > arrayLength {
		numberOfArray = arrayLength
	}

	elementsPerArray := arrayLength / numberOfArray

	result := make([][]int64, numberOfArray)

	for i := 0; i < numberOfArray; i++ {
		startIndex := i * elementsPerArray
		endIndex := (i + 1) * elementsPerArray

		if i == numberOfArray-1 {
			endIndex = arrayLength
		}

		result[i] = arrayInput[startIndex:endIndex]
	}

	return result
}
