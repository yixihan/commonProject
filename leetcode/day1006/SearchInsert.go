package main

func searchInsert(nums []int, target int) int {
	length := len(nums)
	l, r := 0, length-1

	for l < r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
