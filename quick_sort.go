package main

import (
	"fmt"
	"testing"
)

func QuickSort(arr []string) []string {
	pivotIdx := len(arr) - 1
	k := 0
	if len(arr) < 2 {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		if Compare([]byte(arr[i]), []byte(arr[pivotIdx])) == 1 || pivotIdx == i {
			arr[k], arr[i] = arr[i], arr[k]
			if pivotIdx != i {
				k++
			}
		}
	}
	QuickSort(arr[:k])
	QuickSort(arr[k+1:])
	return arr
}

func Test_qs_sort_words_1(t *testing.T) {

	arr := []string{"araba", "kasap", "Zalgiris",
		"Alexander", "The", "abc",
		"aBc", "Patates", "saksı",
		"canta", "ABAH", "tepsi"}

	result := QuickSort(arr)
	if result[0] != "ABAH" ||
		result[1] != "Alexander" ||
		result[2] != "Patates" ||
		result[3] != "The" ||
		result[4] != "Zalgiris" ||
		result[5] != "aBc" ||
		result[6] != "abc" ||
		result[7] != "araba" ||
		result[8] != "canta" ||
		result[9] != "kasap" ||
		result[10] != "saksı" ||
		result[11] != "tepsi" {
		panic(fmt.Sprintf("words not sorted correctly by quick sort."))
	}
}
