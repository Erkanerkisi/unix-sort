package main

import "math"

func MergeSort(arr []string) []string {
	arrSize := len(arr)
	var arrMid = int(math.Floor(float64(arrSize / 2)))

	left := arr[:arrMid]
	right := arr[arrMid:]

	if len(right) > 1 {
		right = MergeSort(right)
	}

	if len(left) > 1 {
		left = MergeSort(left)
	}
	return merge(left, right)
}

func merge(arr1 []string, arr2 []string) []string {
	newArr := make([]string, 0)
	i := 0
	j := 0
	arr1Size := len(arr1)
	arr2Size := len(arr2)
	for i < arr1Size || j < arr2Size {
		if i < arr1Size && j < arr2Size {
			compareResult := Compare([]byte(arr1[i]), []byte(arr2[j]))
			if compareResult == 1 {
				newArr = append(newArr, arr1[i])
				i++
			} else {
				newArr = append(newArr, arr2[j])
				j++
			}
		} else if i < arr1Size {
			for ; i < arr1Size; i++ {
				newArr = append(newArr, arr1[i])
			}
		} else if j < arr2Size {
			for ; j < arr2Size; j++ {
				newArr = append(newArr, arr2[j])
			}
		}
	}
	return newArr
}
