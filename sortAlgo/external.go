package sortAlgo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

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
		fileNameId := strconv.Itoa(i)
		path := common.PATH_TEMP + "/chunk_" + fileNameId + ".bin"

		// Open output files in write mode.
		file := openFile(path, "w")
		defer file.Close()
		out = append(out, file)
	}

	// Allocate a dynamic array large enough to accommodate runs of size common.CHUCK_SIZE

	moreInput := true
	nextOutputFile := 0

	readFile := bufio.NewReader(in)
	for moreInput && nextOutputFile != common.NUMBER_OF_CHUCKS_FILE {
		// Write common.CHUCK_SIZE elements into arr from the input file
		arr := []int64{}

		for i := 0; i < common.CHUNK_SIZE; i++ {
			element, err := readInt64(readFile)
			if err != nil {
				fmt.Println("Read to create chunks fail, err: ", err)
				moreInput = false
				break
			}
			if element == 0 {
				break
			}
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
			var bin_buff bytes.Buffer
			binary.Write(&bin_buff, binary.BigEndian, arr[j])
			_, err := out[nextOutputFile].Write(bin_buff.Bytes())
			if err != nil {
				fmt.Println("Error while write to binary file, err: ", err)
				return nil, err
			}
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
	for ; i < common.NUMBER_OF_CHUCKS_FILE; i++ {
		var element int64
		data := make([]byte, 8)
		_, err := in[i].Read(data)
		if err != nil {
			fmt.Println("Read From binary file fail, err: ", err)
			return err
		}

		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &element)
		if err != nil {
			fmt.Println("Read binary fail, err: ", err)
			return err
		}
		// _, err := fmt.Fscanf(in[i], "%d", &harr[i].Element)
		harr[i].Element = element
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
		var element int64
		data := make([]byte, 8)
		_, err := in[root.I].Read(data)
		if err != nil {
			root.Element = int64(common.MAX_INT)
			count++
		}

		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &element)
		if err != nil {
			fmt.Println("Read binary fail, err: ", err)
			return err
		}

		// if _, err := fmt.Fscanf(in[root.I], "%d", &root.Element); err != nil {
		// 	root.Element = int64(common.MAX_INT) // INT_MAX
		// 	count++
		// }

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

func readInt64(in *bufio.Reader) (int64, error) {
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
