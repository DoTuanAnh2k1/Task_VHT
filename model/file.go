package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"main.go/common"
	"main.go/helper"
)

// File Manager: Read, Write

type FileManager struct{}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) ReadFromFile(filePath string) ([]int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int64
	readFile := bufio.NewReader(file)
	for i := 0; i < common.NUMBER_OF_NUMBER; i++ {
		element, err := helper.ReadInt64(readFile)
		if err != nil {
			fmt.Println("Cant readfrom file, err: ", err)
			return numbers, err
		}
		if element == 0 {
			break
		}
		numbers = append(numbers, element)
	}

	return numbers, nil
}

func (fm *FileManager) WriteToFile(numbers []int64, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, num := range numbers {
		_, err := fmt.Fprintln(file, num)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateData(path string, minValue, maxValue, numCount int64) error {
	t := NewTimer()
	t.Start()
	fmt.Println("===================================================================")
	fmt.Println("Create data")
	fmt.Println("===================================================================")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := numCount; i > 0; i-- {
		number := strconv.FormatInt(i, 10)
		_, err := writer.WriteString(number)
		if err != nil {
			fmt.Println("An cuc roi, error: ", err)
			return err
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer:", err)
			return err
		}
	}
	fmt.Println("Time gen data: ", t.Stop())
	return nil
}
