package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	flag := false

	for i := 1; i < len(arr)-1; i++ {
		if flag = helper(arr, i); flag {
			break
		}
	}

	return flag
}

/*
	在 0 < i < arr.length - 1 条件下，存在 i 使得：
	arr[0] < arr[1] < ... arr[i-1] < arr[i]
	arr[i] > arr[i+1] > ... > arr[arr.length - 1]
*/
func helper(arr []int, index int) bool {
	for i := 0; i < index; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}

	for i := index; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}

	return true
}
