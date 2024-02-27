package sortAlgo

import (
	"bufio"
	"bytes"
	"container/heap"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"main.go/common"
	"main.go/model"
)

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
	total_time_sorting := float64(0)
	total_time_read_file := float64(0)
	total_time_write_file := float64(0)

	for moreInput && nextOutputFile != common.NUMBER_OF_CHUCKS_FILE {
		if nextOutputFile%500 == 0 {
			fmt.Println("Working on file chunk id: ", nextOutputFile)
		}
		// Write common.CHUCK_SIZE elements into arr from the input file
		arr := []int64{}

		timeReadFile := model.NewTimer()
		timeReadFile.Start()
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
		total_time_read_file = total_time_read_file + timeReadFile.Stop()

		time_sorting := model.NewTimer()
		time_sorting.Start()
		// Sort array using library
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		total_time_sorting = total_time_sorting + float64(time_sorting.Stop())

		// Write the records to the appropriate scratch output file
		// Can't assume that the loop runs to runSize
		// Since the last run's length may be less than runSize
		data := []byte{}
		timeWriteFile := model.NewTimer()
		timeWriteFile.Start()

		for j := 0; j < len(arr); j++ {
			var binBuff bytes.Buffer
			binary.Write(&binBuff, binary.BigEndian, arr[j])
			data = append(data, binBuff.Bytes()...)
		}

		_, err := out[nextOutputFile].Write(data)
		if err != nil {
			fmt.Println("Error while write to binary file, err: ", err)
			return nil, err
		}
		total_time_write_file = total_time_write_file + timeWriteFile.Stop()

		nextOutputFile++
	}

	fmt.Println("Total time sorting: 	", total_time_sorting)
	fmt.Println("Total time read file: 	", total_time_read_file)
	fmt.Println("Total time write file: ", total_time_write_file)
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

	writer := bufio.NewWriter(out)
	time_create_pq := model.NewTimer()
	time_create_pq.Start()
	fmt.Println("===================================================================")
	fmt.Println("Create Priority Queue")
	fmt.Println("===================================================================")

	pq := make(model.PriorityQueue, 0)
	data := make([]byte, common.BYTES_BUFF_FILE)
	for i := 0; i < common.NUMBER_OF_CHUCKS_FILE; i++ {
		numberData, err := in[i].Read(data)
		if err != nil {
			fmt.Println("Read From binary file fail, err: ", err)
			return err
		}

		for j := 0; j < numberData; j = j + 8 {
			var element int64
			buffer := bytes.NewBuffer(data[j : j+8])
			err = binary.Read(buffer, binary.BigEndian, &element)
			if err != nil {
				fmt.Println("Read from buffer fail, err: ", err)
				return err
			}
			heap.Push(&pq, &model.Item{
				FileId:   i,
				Priority: element,
			})
		}

		if numberData < common.BYTES_BUFF_FILE {
			in[i].Close()
		}
	}
	fmt.Println("Time init pq: ", time_create_pq.Stop())

	time_merge_chunks := model.NewTimer()
	time_merge_chunks.Start()
	checkRemain := make([]int, common.NUMBER_OF_CHUCKS_FILE)
	bufferAnswer := ""
	countBuffer := 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*model.Item)
		bufferAnswer = bufferAnswer + strconv.FormatInt(item.Priority, 10) + "\n"
		countBuffer++

		if countBuffer == common.COUNT_BUFFER {
			// fmt.Fprint(out, bufferAnswer)
			writer.WriteString(bufferAnswer)
			bufferAnswer = ""
			countBuffer = 0
		}

		checkRemain[item.FileId]++
		if checkRemain[item.FileId] == common.BYTES_BUFF_FILE/8 {
			checkRemain[item.FileId] = 0

			numberData, err := in[item.FileId].Read(data)
			if err != nil {
				fmt.Println("Read From binary file fail, err: ", err)
				in[item.FileId].Close()
			}

			for i := 0; i < numberData; i = i + 8 {
				var element int64
				buffer := bytes.NewBuffer(data[i : i+8])
				err = binary.Read(buffer, binary.BigEndian, &element)
				if err != nil {
					fmt.Println("Read from buffer fail, err: ", err)
					return err
				}
				heap.Push(&pq, &model.Item{
					FileId:   item.FileId,
					Priority: element,
				})
			}

			if numberData < common.BYTES_BUFF_FILE {
				in[item.FileId].Close()
			}
		}
	}

	if len(bufferAnswer) != 0 {
		// fmt.Fprint(out, bufferAnswer)
		writer.WriteString(bufferAnswer)
	}

	fmt.Println("Time merge file: 			", time_merge_chunks.Stop())
	return nil
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
