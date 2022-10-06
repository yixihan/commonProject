package main

func removeDuplicates(nums []int) int {
	l, r := 0, 1

	for l < r && r < len(nums) {
		if nums[l] != nums[r] {
			nums[l+1] = nums[r]
			l++
		}
		r++
	}

	return l + 1
}
