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

	f := model.NewFileManager()
	flag := true
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

	arrayInPut, err := f.ReadFromFile(common.PATH_INPUT)
	if err != nil {
		fmt.Println("Cannot Read Input File: ", err.Error())
		return
	}

	// Add goroutine running merge sort algorithm without separate array
	wg.Add(1)
	go worker.WorkerMergeSortOnly(arrayInPut[:], &wg)

	// Add goroutine running merge sort algorithm with separate array
	wg.Add(1)
	go worker.WorkerMergeSortMulti(arrayInPut[:], &wg)

	wg.Add(1)
	go worker.WorkerLibSort(arrayInPut[:], &wg)

	wg.Add(1)
	go worker.WorkerMergeSortExternal(&wg)
	wg.Wait()
}
