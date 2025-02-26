package arrays

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		if nums[i] != nums[j] {
			// swap - means it will push 0 to the right
			if nums[i] == 0 {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
			i++
			j++
			continue
		}
		j++
	}

	//fmt.Printf("Final nums is %+v", nums)
}
