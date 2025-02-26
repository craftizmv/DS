package arrays

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

func SingleNumber(nums []int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor = nums[i] ^ xor
	}

	return xor
}

//
//
// [4,1,2,1,2] = 100, 001, 010, 001, 010
// 101, 010, 001, 010
// 111, 001, 010
// 110, 010
// 101

// 111 011 100

// [2,3,2],
// 10 11 = 01 10 = 11
