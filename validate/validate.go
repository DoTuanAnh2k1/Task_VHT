package validate

import (
	"fmt"

	"main.go/model"
)

func Validate(filePath string) (bool, error) {
	f := model.NewFileManager()
	array, err := f.ReadFromFile(filePath)
	if err != nil {
		fmt.Println("Validate, Cannot Read File: ", err)
		return false, err
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false, nil
		}
	}

	return true, nil
}
