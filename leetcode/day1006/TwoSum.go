package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		val, ok := m[num]
		if !ok {
			m[target-num] = i
		} else {
			return []int{i, val}

		}
	}
	return []int{}
}
