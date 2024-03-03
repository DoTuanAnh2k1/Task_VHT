package validate

import (
	"bufio"
	"fmt"

	"main.go/common"
	"main.go/helper"
)

func Validate(filePath string) (bool, error) {
	ouputFile := helper.OpenFile(filePath, "r")
	readFile := bufio.NewReader(ouputFile)
	tmp, err := helper.ReadInt64(readFile)
	if err != nil {
		fmt.Println("Read to validate file fail, error: ", err)
		return false, err
	}
	for i := 0; i < common.NUMBER_OF_NUMBER-1; i++ {
		element, err := helper.ReadInt64(readFile)
		if err != nil {
			fmt.Println("Read to validate file fail, error: ", err)
			return false, err
		}
		if element == 0 {
			break
		}
		if tmp > element {
			return false, nil
		}
		tmp = element
	}

	return true, nil
}
