package main

func BubbleSort(arr []string) []string {
	s := true
	k := len(arr) - 1
	for s {
		s, k = swap(arr, k)
	}
	return arr
}

func swap(arr []string, k int) (bool, int) {
	swapped := false
	for i := 0; i < k; i++ {
		if Compare([]byte(arr[i]), []byte(arr[i+1])) == 2 {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
			swapped = true
		}
	}
	return swapped, k - 1
}
