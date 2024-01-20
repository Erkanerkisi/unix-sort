package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var unique bool
	var headItemSize int
	flag.BoolVar(&unique, "u", false, "unique data")
	flag.IntVar(&headItemSize, "head", 5, "how many head data we are gonna show")

	flag.Parse()
	args := flag.Args()

	if args[0] == "" {
		panic("you need to pass file name.")
	}
	fileName := args[0]

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("can not open file")
	}

	uniqueMap := make(map[string]bool)
	fmt.Println(unique)
	str := strings.Builder{}
	words := make([]string, 0)
	for i := 0; i < len(file); i++ {
		if file[i] == 10 {

			if !unique {
				words = append(words, str.String())
			} else {
				if _, ok := uniqueMap[str.String()]; !ok {
					words = append(words, str.String())
					uniqueMap[str.String()] = true
				}
			}
			str.Reset()
			continue
		}
		str.WriteByte(file[i])
	}

	result := QuickSort(words)
	// printing result
	for i := 0; i < headItemSize; i++ {
		fmt.Println(result[i])
	}
}
