package day1008

import (
	"sort"
)

func advantageCount(nums1 []int, nums2 []int) []int {
	n, minIdx := len(nums1), 0
	visited := make([]bool, n, n)
	ans := make([]int, 0, n)
	sort.Ints(nums1)

	for _, val := range nums2 {
		l, r := 0, n

		for l != r {
			mid := l + ((r - l) >> 1)
			if nums1[mid] < val {
				l = mid + 1
			} else {
				r = mid
			}
		}

		for l < n && (visited[l] || nums1[l] == val) {
			l++
		}

		if l == n {
			for visited[minIdx] {
				minIdx++
			}
			ans = append(ans, nums1[minIdx])
			visited[minIdx] = true
			minIdx++
		} else {
			visited[l] = true
			ans = append(ans, nums1[l])
		}
	}
	return ans
}
