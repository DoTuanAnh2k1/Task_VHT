package main

import (
	"fmt"
	"sync"

	"main.go/common"
	"main.go/model"
	"main.go/worker"
)

func main() {
	var wg sync.WaitGroup
	// Create Data Input
	flag := false
	if flag {
		err := model.CreateData(
			common.PATH_INPUT,
			common.MINVALUE,
			common.MAXVALUE,
			common.NUMBER_OF_NUMBER,
		)
		if err != nil {
			fmt.Println("Cannot Create Data: ", err.Error())
			return
		}
	}
	// f := model.NewFileManager()

	// arrayInPut, err := f.ReadFromFile(common.PATH_INPUT)
	// if err != nil {
	// 	fmt.Println("Cannot Read Input File: ", err.Error())
	// 	return
	// }

	// // Add goroutine running merge sort algorithm without separate array
	// wg.Add(1)
	// arrayInPutMergeSortOnly := make([]int64, common.NUMBER_OF_NUMBER)
	// copy(arrayInPutMergeSortOnly, arrayInPut)
	// go worker.WorkerMergeSortOnly(arrayInPutMergeSortOnly, &wg)

	// // Add goroutine running merge sort algorithm with separate array
	// wg.Add(1)
	// arrayInPutMergeMultiOnly := make([]int64, common.NUMBER_OF_NUMBER)
	// copy(arrayInPutMergeMultiOnly, arrayInPut)
	// go worker.WorkerMergeSortMulti(arrayInPutMergeMultiOnly, &wg)

	// wg.Add(1)
	// arrayInPutLib := make([]int64, common.NUMBER_OF_NUMBER)
	// copy(arrayInPutLib, arrayInPut)
	// go worker.WorkerLibSort(arrayInPutLib, &wg)

	wg.Add(1)
	go worker.WorkerMergeSortExternal(&wg)

	wg.Wait()
}
