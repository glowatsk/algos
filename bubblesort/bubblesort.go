package bubble

// 0n2
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums = sweep(nums)
	}
	return nums
}

func sweep(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		first := n[i]
		second := n[i+1]
		if first > second {
			n[i], n[i+1] = n[i+1], n[i]
		}
	}
	return n
}
