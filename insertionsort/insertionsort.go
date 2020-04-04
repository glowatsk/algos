package insertion

// should be n complexity
func insertionSort(list []int) []int {
	var sorted []int
	for _, item := range list {
		sorted = sweep(sorted, item)
	}
	for i, val := range sorted {
		list[i] = val
	}
	return list
}

func sweep(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}
