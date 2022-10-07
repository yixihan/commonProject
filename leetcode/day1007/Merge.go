package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := copyArray(nums1, m)
	index := 0
	i, j := 0, 0

	for i < m && j < n {
		if tmp[i] <= nums2[j] {
			nums1[index] = tmp[i]
			i++
		} else {
			nums1[index] = nums2[j]
			j++
		}
		index++
	}

	for i < m {
		nums1[index] = tmp[i]
		i++
		index++
	}

	for j < n {
		nums1[index] = nums2[j]
		j++
		index++
	}
}

func copyArray(a []int, m int) []int {
	arr := make([]int, m)
	index := 0
	for i := 0; i < m; i++ {
		arr[index] = a[i]
		index++
	}
	return arr
}
