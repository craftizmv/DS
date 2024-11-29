package arrays

// Remove duplicates in sorted array

func removeDuplicates(nums []int) int {
	// pointer 1
	idx := 0

	// i index is acting as a second pointer.
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			idx = idx + 1
			nums[idx] = nums[i]
		}
	}

	return idx + 1
}
