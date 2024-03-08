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

	// // Add goroutine running merge sort algorithm without separate array
	// wg.Add(1)
	// go worker.WorkerMergeSortOnly(&wg)

	// // Add goroutine running merge sort algorithm with separate array
	// wg.Add(1)
	// go worker.WorkerMergeSortMulti(&wg)

	// wg.Add(1)
	// go worker.WorkerLibSort(&wg)

	wg.Add(1)
	go worker.WorkerMergeSortExternal(&wg)

	wg.Wait()
}
