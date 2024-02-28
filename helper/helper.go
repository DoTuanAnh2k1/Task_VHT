package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Separate an array to numberOfArray arrays smaller
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

// Read Int64 fast from file, used to use in competitive programming with GoLang
func ReadInt64(in *bufio.Reader) (int64, error) {
	nStr, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		fmt.Println("Error reading file, error: ", err)
		return 0, err
	}
	return n, nil
}

// OpenFile opens a file with the given name and mode
func OpenFile(fileName, mode string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	return file
}
