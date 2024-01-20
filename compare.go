package main

// 1 => first
// 2 => second
// 0 => same
func Compare(first, second []byte) int {
	firstSize := len(first)
	secondSize := len(second)
	if firstSize < 1 || secondSize < 1 {
		panic("first or second word is empty.")
	}

	minSize := firstSize

	if firstSize > secondSize {
		minSize = secondSize
	}

	for i := 0; i < minSize; i++ {
		if first[i] < second[i] {
			return 1
		} else if first[i] > second[i] {
			return 2
		}
	}
	if firstSize < secondSize {
		return 1
	} else if secondSize < firstSize {
		return 2
	}
	return 0
}
