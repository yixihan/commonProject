package main

func maxAscendingSum(nums []int) int {
	sum, maxSum := nums[0], 0

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			maxSum = max(maxSum, sum)
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}

	return max(sum, maxSum)
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
