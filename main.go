package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const newLine byte = 10

const (
	QuickSortType = iota
	BubbleSortType
	MergeSortType
	RandomSortType
)

func main() {
	var unique bool
	var headItemSize int
	var sortType int

	flag.BoolVar(&unique, "u", false, "unique data")
	flag.IntVar(&headItemSize, "head", 5, "how many head data we are gonna show")
	flag.IntVar(&sortType, "sort", QuickSortType, "sorting type (0: quicksort, 1: bubblesort, 2: mergesort, 3: randomsort)")

	flag.Parse()
	args := flag.Args()

	if args[0] == "" {
		log.Fatal("You need to pass a file name.")
	}
	fileName := args[0]

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Cannot open file:", err)
	}

	words := parse(file, unique)

	sortingFunctions := map[int]func([]string) []string{
		QuickSortType:  QuickSort,
		BubbleSortType: BubbleSort,
		MergeSortType:  MergeSort,
		RandomSortType: RandomSort,
	}

	result := sortingFunctions[sortType](words)

	// printing result
	for i := 0; i < headItemSize && i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func parse(file []byte, unique bool) []string {
	uniqueMap := make(map[string]bool)
	str := strings.Builder{}
	words := make([]string, 0)
	for i := 0; i < len(file); i++ {
		if file[i] == newLine {
			if !unique || unique && !uniqueMap[str.String()] {
				words = append(words, str.String())
				if unique {
					uniqueMap[str.String()] = true
				}
			}
			str.Reset()
			continue
		}
		str.WriteByte(file[i])
	}
	return words
}
