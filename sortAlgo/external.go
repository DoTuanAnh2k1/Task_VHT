package sortAlgo

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"main.go/common"
	"main.go/model"
)

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

		chunk := []int64{}
		for i := 0; i < common.CHUNK_SIZE && scanner.Scan(); i++ {
			valueText := scanner.Text()
			value, err := strconv.ParseInt(valueText, 10, 64)
			if err != nil {
				return nil, err
			}
			chunk = append(chunk, value)
		}

		sort.Slice(chunk, func(i, j int) bool {
			return chunk[i] < chunk[j]
		})

		for _, value := range chunk {
			line := strconv.FormatInt(value, 10)
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

func MergeChunks(chunkFileNames []string, outputFilePath string) error {
	var in []*os.File

	// Open input files
	for i := 0; i < len(chunkFileNames); i++ {
		fileName := strconv.Itoa(i)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return err
		}
		defer file.Close()
		in = append(in, file)
	}

	// Open output file
	out, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}
	defer out.Close()

	// Create a min heap with k heap nodes
	harr := make([]model.MinHeapNode, len(outputFilePath))
	i := 0
	for ; i < len(chunkFileNames); i++ {
		// Break if no output file is empty and index i will be the number of input files
		if _, err := fmt.Fscanf(in[i], "%d", &harr[i].Element); err != nil {
			break
		}

		// Index of scratch output file
		harr[i].I = i
	}

	// Create the heap
	hp := model.NewMinHeap(harr[:i], i)
	count := 0

	// Now one by one get the minimum element from the min heap and replace it with the next element
	// Run until all filled input files reach EOF
	for count != i {
		// Get the minimum element and store it in the output file
		root := hp.GetMin()
		fmt.Fprintf(out, "%d ", root.Element)

		// Find the next element that will replace the current root of the heap.
		// The next element belongs to the same input file as the current min element.
		if _, err := fmt.Fscanf(in[root.I], "%d", &root.Element); err != nil {
			root.Element = int(^uint(0) >> 1) // INT_MAX
			count++
		}

		// Replace root with the next element of the input file
		hp.ReplaceMin(root)
	}
	return nil
}

// OpenFile opens a file with the given name and mode
func OpenFile(fileName, mode string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	return file, err
}

// External Merge Sort algorithm

func ExternalMergeSort(inputFilePath, outputFilePath string) error {
	chunkFileNames, err := CreateChunks(inputFilePath)
	if err != nil {
		return err
	}

	err = MergeChunks(chunkFileNames, outputFilePath)
	if err != nil {
		return err
	}

	for _, chunkFileName := range chunkFileNames {
		if err := os.Remove(chunkFileName); err != nil {
			return err
		}
	}

	return nil
}
