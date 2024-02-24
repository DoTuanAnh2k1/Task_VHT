package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	for i := int64(0); i < numCount; i++ {
		number := rand.Int63n(maxValue-minValue+1) + minValue
		_, err := fmt.Fprintln(file, number)
		if err != nil {
			return err
		}
	}
	fmt.Println("Time gen data: ", t.Stop())
	return nil
}
