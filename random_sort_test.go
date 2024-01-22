package main

import (
	"fmt"
	"testing"
)

func Test_RandomSort(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result := RandomSort(arr)
	if result[0] != "1" &&
		result[1] != "2" &&
		result[2] != "3" &&
		result[3] != "4" &&
		result[4] != "5" &&
		result[5] != "6" &&
		result[6] != "7" &&
		result[7] != "8" &&
		result[8] != "9" {
		panic(fmt.Sprintf("words not sorted correctly by quick sort."))
	}
	fmt.Println(arr)
}
