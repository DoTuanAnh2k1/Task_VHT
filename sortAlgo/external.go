package sortAlgo

import (
	"bufio"
	"bytes"
	"container/heap"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"sync"

	"main.go/common"
	"main.go/helper"
	"main.go/model"
)

func workerSort(arr []int64, outputFile *os.File, wg *sync.WaitGroup) error {
	arrSep := helper.SeparateArray(common.NUMBER_OF_GOROUTINE, arr)
	var wgChild sync.WaitGroup
	for i := 0; i < common.NUMBER_OF_GOROUTINE; i++ {
		wgChild.Add(1)
		go SortingMulti(arrSep[i], &wgChild)
	}
	wgChild.Wait()
	arr = MergeMultiArray(arrSep)

	data := []byte{}

	for j := 0; j < len(arr); j++ {
		var binBuff bytes.Buffer
		binary.Write(&binBuff, binary.BigEndian, arr[j])
		data = append(data, binBuff.Bytes()...)
	}
	_, err := outputFile.Write(data)
	if err != nil {
		fmt.Println("Error while write to binary file, err: ", err)
		return err
	}
	outputFile.Close()
	wg.Done()
	return nil
}

func CreateChunks(inputFilePath string) ([]*os.File, error) {
	// For big input file
	in := helper.OpenFile(inputFilePath, "r")
	defer in.Close()

	// Output chunks files
	var out []*os.File
	for i := 0; i < common.NUMBER_OF_CHUCKS_FILE; i++ {
		// Convert i to string
		fileNameId := strconv.Itoa(i)
		path := common.PATH_TEMP + "/chunk_" + fileNameId + ".bin"

		// Open output files in write mode.
		file := helper.OpenFile(path, "w")
		defer file.Close()
		out = append(out, file)
	}

	// Read from input file
	moreInput := true
	nextOutputFile := -1
	readFile := bufio.NewReader(in)
	processNums := 0
	totalTimeReadFile := float64(0)
	var wgSort sync.WaitGroup

	for moreInput && nextOutputFile != common.NUMBER_OF_CHUCKS_FILE {
		nextOutputFile++
		if processNums%(common.NUMBER_OF_NUMBER/10) == 0 {
			fmt.Println("Processing: ", (processNums*100)/common.NUMBER_OF_NUMBER, " %, time: ", totalTimeReadFile)
		}

		// Create an array to store and sort number in input file
		if nextOutputFile == common.NUMBER_OF_CHUCKS_FILE {
			break
		}
		arr := []int64{}

		timeReadFile := model.NewTimer()
		timeReadFile.Start()
		// Write common.CHUCK_SIZE elements into arr from the input file
		for i := 0; i < common.CHUNK_SIZE; i++ {
			element, err := helper.ReadInt64(readFile)
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
		totalTimeReadFile = totalTimeReadFile + timeReadFile.Stop()
		processNums = processNums + len(arr)
		wgSort.Add(1)
		go workerSort(arr, out[nextOutputFile], &wgSort)
	}
	wgSort.Wait()

	fmt.Println("Total time read file: 	", totalTimeReadFile)
	return out, nil
}

func MergeChunks(out_createChunks []*os.File, outputFilePath string) error {
	// Open chunks files
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

	// Create and open output file
	out, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}
	defer out.Close()

	// Create Priority Queue and
	// pushing common.BYTES_BUFF_FILE / 8 numbers
	// per chunk files to PQ
	writer := bufio.NewWriter(out)
	fmt.Println("Create Priority Queue")

	//* Create object pool here
	poolObj := model.NewObjectPool(common.POOL_SIZE)

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

			acquire := poolObj.Acquire(i, element)
			heap.Push(&pq, acquire)
		}

		// Case that if number of byte read from file
		// less than byte buffer then close the file
		// because file have nothing left
		if numberData < common.BYTES_BUFF_FILE {
			in[i].Close()
		}
	}

	timeMergeChunks := model.NewTimer()
	timeMergeChunks.Start()
	checkRemain := make([]int, common.NUMBER_OF_CHUCKS_FILE)
	bufferAnswer := ""
	countBuffer := 0
	processNums := 0

	// While len PQ > 0, push the top of queue.
	// If that element of file remain == 0 then push more
	// number of that file in queue.
	// When count buffer == common.COUNT_BUFFER then write bufferAns
	// to output file and restart the count.
	for pq.Len() > 0 {
		if processNums%(common.NUMBER_OF_NUMBER/10) == 0 {
			fmt.Println("Processing ", (processNums*100)/common.NUMBER_OF_NUMBER, "%")
		}
		item := heap.Pop(&pq).(*model.Item)
		bufferAnswer = bufferAnswer + strconv.FormatInt(item.Priority, 10) + "\n"
		countBuffer++
		processNums++

		if countBuffer == common.COUNT_BUFFER {
			_, err := writer.WriteString(bufferAnswer)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return err
			}
			err = writer.Flush()
			if err != nil {
				fmt.Println("Error flushing buffer:", err)
				return err
			}
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

				heap.Push(
					&pq,
					poolObj.Acquire(
						item.FileId,
						element,
					),
				)
			}

			if numberData < common.BYTES_BUFF_FILE {
				in[item.FileId].Close()
			}
		}

		poolObj.Release(item)
	}

	// In case that pq len empty but the buffer Answer
	// remain haven't been written to output file
	if len(bufferAnswer) != 0 {
		_, err := writer.WriteString(bufferAnswer)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer:", err)
			return err
		}
	}

	fmt.Println("Time merge file: 			", timeMergeChunks.Stop())
	return nil
}

// External Merge Sort algorithm
func ExternalMergeSort(inputFilePath, outputFilePath string) error {
	// Create chunks phrase
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

	// Merge chunks phrase
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

	// Remove chunks
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
