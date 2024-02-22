package helper

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"main.go/common"
)

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

/*
	Input:
		- inputFilePath: path to file contains input 		Type: string
	Output:
		- List of file created								Type: []string
		- error
*/

func CreateChunks(inputFilePath string) ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	chunkFileNames := []string{}

	scanner := bufio.NewScanner(file)
	chunkIndex := 0
	for scanner.Scan() {
		path_chunk := common.PATH_TEMP + "/chunk_%d.txt"
		chunkFileName := fmt.Sprintf(path_chunk, chunkIndex)
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			return nil, err
		}
		chunkFileNames = append(chunkFileNames, chunkFileName)

		chunk := []string{}
		for i := 0; i < common.CHUNK_SIZE && scanner.Scan(); i++ {
			chunk = append(chunk, scanner.Text())
		}

		sort.Strings(chunk)
		for _, line := range chunk {
			fmt.Fprintln(chunkFile, line)
		}

		chunkFile.Close()
		chunkIndex++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return chunkFileNames, nil
}
