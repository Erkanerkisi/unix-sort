package main

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
