package worker

import (
	"fmt"
	"sync"

	"main.go/common"
	"main.go/helper"
	"main.go/model"
	"main.go/sortAlgo"
	"main.go/validate"
)

func WorkerMergeSortMulti(arrayInput []int64, wg *sync.WaitGroup) {
	t := model.NewTimer()
	var wgChild sync.WaitGroup
	f := model.NewFileManager()
	t.Start()
	// Separate Array
	array_separate := helper.SeparateArray(common.NUMBER_OF_GOROUTINE, arrayInput)
	for i := 0; i < common.NUMBER_OF_GOROUTINE; i++ {
		wgChild.Add(1)
		go sortAlgo.SortingMulti(array_separate[i], &wgChild)
	}

	wgChild.Wait()
	runtime := t.Stop()
	// exec Multi Merge Sort algorithm
	arrayOutput := sortAlgo.MergeMultiArray(array_separate)
	err := f.WriteToFile(arrayOutput, common.PATH_OUTPUT_MERGESORT_MULTI)
	if err != nil {
		fmt.Println("Cannot Write to file: ", err.Error())
		return
	}

	// validate solution
	validate, err := validate.Validate(common.PATH_OUTPUT_MERGESORT_MULTI)
	if err != nil {
		fmt.Println("Cannot Validate: ", err.Error())
		return
	}

	if validate {
		fmt.Println("Sorting Merge Sort Multi complete, runtime: ", runtime, "second")
	} else {
		fmt.Println("khong on roi dai vuong oi, runtime: ", runtime)
	}
	wg.Done()
}
