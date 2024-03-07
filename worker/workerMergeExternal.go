package worker

import (
	"fmt"
	"sync"

	"main.go/common"
	"main.go/model"
	"main.go/sortAlgo"
)

func WorkerMergeSortExternal(wg *sync.WaitGroup) {
	t := model.NewTimer()

	t.Start()
	// exec External Merge Sort algorithm
	err := sortAlgo.ExternalMergeSort(common.PATH_INPUT, common.PATH_OUTPUT_MERGESORT_EXTERNAL)
	if err != nil {
		fmt.Println("Error Merge External Sort: ", err.Error())
		return
	}
	runtime := t.Stop()
	fmt.Println("Sorting Merge Sort External complete, runtime: ", runtime)

	wg.Done()
}
