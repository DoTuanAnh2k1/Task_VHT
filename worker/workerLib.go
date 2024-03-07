package worker

import (
	"fmt"
	"sync"

	"main.go/common"
	"main.go/model"
	"main.go/sortAlgo"
)

func WorkerLibSort(arr []int64, wg *sync.WaitGroup) {
	t := model.NewTimer()
	f := model.NewFileManager()
	t.Start()

	arrayOutput := sortAlgo.LibSort(arr)
	runtime := t.Stop()

	err := f.WriteToFile(arrayOutput, common.PATH_OUTPUT_LIB_SORT)
	if err != nil {
		fmt.Println("Cannot Write to file Merge Sort Only: ", err.Error())
	}
	fmt.Println("Sorting Lib Sort Only complete, runtime: ", runtime)

	// // validate solution
	// validate, err := validate.Validate(common.PATH_OUTPUT_LIB_SORT)
	// if err != nil {
	// 	fmt.Println("Cannot Validate Lib Sort Only: ", err.Error())
	// 	return
	// }

	// if validate {
	// 	fmt.Println("Sorting Lib Sort Only complete, runtime: ", runtime)
	// } else {
	// 	fmt.Println("Sorting Lib Sort Only Fail, runtime: ", runtime)
	// }
	wg.Done()
}