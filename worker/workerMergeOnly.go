package worker

import (
	"fmt"
	"sync"

	"main.go/common"
	"main.go/model"
	"main.go/sortAlgo"
)

func WorkerMergeSortOnly(arr []int64, wg *sync.WaitGroup) {
	t := model.NewTimer()
	f := model.NewFileManager()
	t.Start()

	// exec Merge Sort algorithm
	arrayOutput := sortAlgo.MergeSort(arr)
	runtime := t.Stop()

	err := f.WriteToFile(arrayOutput, common.PATH_OUTPUT_MERGESORT_ONLY)
	if err != nil {
		fmt.Println("Cannot Write to file Merge Sort Only: ", err.Error())
	}
	fmt.Println("Merge Sort Only complete, runtime: ", runtime)
	// // validate solution
	// validate, err := validate.Validate(common.PATH_OUTPUT_MERGESORT_ONLY)
	// if err != nil {
	// 	fmt.Println("Cannot Validate Merge Sort Only: ", err.Error())
	// 	return
	// }

	// if validate {
	// 	fmt.Println("Sorting Merge Sort Only complete, runtime: ", runtime, "second")
	// } else {
	// 	fmt.Println("Sorting Merge Sort Only Fail")
	// }
	wg.Done()
}
