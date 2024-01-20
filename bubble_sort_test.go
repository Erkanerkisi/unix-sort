package main

import (
	"fmt"
	"testing"
)

func Test_bubble_sort_words_1(t *testing.T) {

	arr := []string{"araba", "kasap", "Zalgiris", "Alexander", "The", "abc", "aBc"}

	result := BubbleSort(arr)
	if result[0] != "Alexander" ||
		result[1] != "The" ||
		result[2] != "Zalgiris" ||
		result[3] != "aBc" ||
		result[4] != "abc" ||
		result[5] != "araba" ||
		result[6] != "kasap" {
		panic(fmt.Sprintf("words not sorted correctly by bubble sort."))
	}
}
