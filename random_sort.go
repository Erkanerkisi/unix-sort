package main

import (
	"math/rand"
)

func RandomSort(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		rdm := rand.Int63n(int64(len(arr)))
		arr[i], arr[rdm] = arr[rdm], arr[i]
	}
	return arr
}
