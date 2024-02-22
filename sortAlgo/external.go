package sortAlgo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"main.go/helper"
	"main.go/model"
)

// External Merge Sort algorithm

func ExternalMergeSort(inputFilePath, outputFilePath string) error {
	chunkFileNames, err := helper.CreateChunks(inputFilePath)
	if err != nil {
		return err
	}

	err = MergeChunks(chunkFileNames, outputFilePath)
	if err != nil {
		return err
	}

	for _, chunkFileName := range chunkFileNames {
		if err := os.Remove(chunkFileName); err != nil {
			return err
		}
	}

	return nil
}

func MergeChunks(chunkFileNames []string, outputFilePath string) error {
	files := make([]*os.File, len(chunkFileNames))
	scanners := make([]*bufio.Scanner, len(chunkFileNames))
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	heap := make([]model.MinHeapNode, len(chunkFileNames))
	for i := range heap {
		files[i], err = os.Open(chunkFileNames[i])
		if err != nil {
			return err
		}
		defer files[i].Close()

		scanners[i] = bufio.NewScanner(files[i])
		if scanners[i].Scan() {
			val, _ := strconv.Atoi(scanners[i].Text())
			heap[i] = model.MinHeapNode{
				Value: val,
				Index: i,
			}
		}
	}

	minHeap := &model.MinHeap{
		Nodes: heap,
	}
	minHeap.Init()

	for minHeap.Len() > 0 {
		node := minHeap.Pop()
		fmt.Fprintln(outputFile, node.Value)

		if scanners[node.Index].Scan() {
			val, _ := strconv.Atoi(scanners[node.Index].Text())
			minHeap.Push(model.MinHeapNode{
				Value: val,
				Index: node.Index,
			})
		}
	}

	return nil
}
