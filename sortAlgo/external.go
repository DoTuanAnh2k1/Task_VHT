package sortAlgo

import (
	"encoding/binary"
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

func CreateChunks(inputFilePath string) ([]*os.File, error) {
	// For big input file
	in := openFile(inputFilePath, "r")
	defer in.Close()

	// Output scratch files
	var out []*os.File
	for i := 0; i < common.NUMBER_OF_CHUCKS_FILE; i++ {
		// Convert i to string
		fileName := strconv.Itoa(i)
		path := common.PATH_TEMP + "/chunk_" + fileName + ".txt"

		// Open output files in write mode.
		file := openFile(path, "w")
		defer file.Close()
		out = append(out, file)
	}

	// Allocate a dynamic array large enough to accommodate runs of size common.CHUCK_SIZE

	moreInput := true
	nextOutputFile := 0

	for moreInput && nextOutputFile != common.NUMBER_OF_CHUCKS_FILE {
		// Write common.CHUCK_SIZE elements into arr from the input file
		arr := []int64{}
		buffer := make([]byte, 8) // 8 bytes for int64

		for i := 0; i < common.CHUNK_SIZE; i++ {
			_, err := in.Read(buffer)
			if err != nil {
				fmt.Println("Read buffer fail, error: ", err)
				moreInput = false
				break
			}

			element := int64(binary.LittleEndian.Uint64(buffer))
			arr = append(arr, element)
		}

		// Sort array using library
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})

		// Write the records to the appropriate scratch output file
		// Can't assume that the loop runs to runSize
		// Since the last run's length may be less than runSize
		for j := 0; j < len(arr); j++ {
			fmt.Fprintf(out[nextOutputFile], "%d ", arr[j])
		}

		nextOutputFile++
	}

	// fmt.Println(nextOutputFile)
	// Close input and output files
	for i := 0; i < common.NUMBER_OF_CHUCKS_FILE; i++ {
		out[i].Close()
	}
	return out, nil
}

// openFile opens a file with the given name and mode
func openFile(fileName, mode string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	return file
}

func MergeChunks(out_createChunks []*os.File, outputFilePath string) error {
	var in []*os.File
	for _, fileOut := range out_createChunks {
		file, err := os.Open(fileOut.Name())
		if err != nil {
			fmt.Println("Error opening file: ", err)
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

	fmt.Println("===================================================================")
	fmt.Println("Create Min Heap")
	fmt.Println("===================================================================")
	// Create a min heap with k heap nodes
	harr := make([]model.MinHeapNode, common.NUMBER_OF_CHUCKS_FILE)
	i := 0
	buffer := make([]byte, 8)
	for ; i < common.NUMBER_OF_CHUCKS_FILE; i++ {
		// Break if no output file is empty and index i will be the number of input files

		// _, err := fmt.Fscanf(in[i], "%d", &harr[i].Element)
		_, err := in[i].Read(buffer)
		// fmt.Println(value)
		if err != nil {
			fmt.Println(err)
			break
		}

		harr[i].Element = int64(binary.LittleEndian.Uint64(buffer))
		// Index of scratch output file
		harr[i].I = i
	}
	// fmt.Println(harr)
	// Create the heap
	hp := model.NewMinHeap(harr[:], i)
	count := 0
	// fmt.Println(hp)
	fmt.Println("===================================================================")
	fmt.Println("Merge it")
	fmt.Println("===================================================================")
	// Now one by one get the minimum element from the min heap and replace it with the next element
	// Run until all filled input files reach EOF
	for count != i {
		// Get the minimum element and store it in the output file
		root := hp.GetMin()
		fmt.Fprintf(out, "%d\n", root.Element)

		// Find the next element that will replace the current root of the heap.
		// The next element belongs to the same input file as the current min element.
		if _, err := fmt.Fscanf(in[root.I], "%d", &root.Element); err != nil {
			root.Element = int64(common.MAX_INT) // INT_MAX
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
	time_create_chunks := model.NewTimer()
	time_create_chunks.Start()
	fmt.Println("===================================================================")
	fmt.Println("Create Chunks")
	fmt.Println("===================================================================")
	chunkFileNames, err := CreateChunks(inputFilePath)
	if err != nil {
		fmt.Println("External Merge Sort, error: ", err)
		return err
	}
	fmt.Println("Create Chunks success, runtime: ", time_create_chunks.Stop())

	time_merge_chunks := model.NewTimer()
	time_merge_chunks.Start()
	fmt.Println("===================================================================")
	fmt.Println("Merge Chunks")
	fmt.Println("===================================================================")

	err = MergeChunks(chunkFileNames, outputFilePath)
	if err != nil {
		return err
	}
	fmt.Println("Merge Chunks success, runtime: ", time_merge_chunks.Stop())

	time_remove_chunks := model.NewTimer()
	time_remove_chunks.Start()
	fmt.Println("===================================================================")
	fmt.Println("Remove Chunks")
	fmt.Println("===================================================================")

	for _, chunkFileName := range chunkFileNames {
		err := os.Remove(chunkFileName.Name())
		if err != nil {
			return err
		}
	}

	fmt.Println("Remove Chunks success, runtime: ", time_remove_chunks.Stop())

	return nil
}
