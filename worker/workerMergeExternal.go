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
	err := sortAlgo.ExternalMergeSort(common.PATH_INPUT, common.PATH_OUTPUT_MERGESORT_ONLY)
	if err != nil {
		fmt.Println("Error Merge External Only: ", err.Error())
	}
	runtime := t.Stop()
	fmt.Println("Sorting Merge Sort External complete, runtime: ", runtime, "second")

	wg.Done()
}
