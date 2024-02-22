package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"main.go/common"
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
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

func generateRandomNumbers(n, min, max int64) []int64 {
	numbers := make([]int64, n)
	rand.Seed(time.Now().UnixNano())

	for i := int64(0); i < n; i++ {
		numbers[i] = rand.Int63n(max-min+1) + min
	}

	return numbers
}

func CreateData(path string, minValue, maxValue, numCount int64) (*FileManager, error) {
	numbers := generateRandomNumbers(numCount, minValue, maxValue)

	file := NewFileManager()

	err := file.WriteToFile(numbers, common.PATH_INPUT)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return file, err
	}

	return file, nil
}
