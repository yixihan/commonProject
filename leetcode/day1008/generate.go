package main

func generate(numRows int) [][]int {
	ans := make([][]int, numRows, numRows)

	for i := 1; i <= numRows; i++ {
		ans[i-1] = make([]int, i, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				ans[i-1][j] = 1
			} else {
				ans[i-1][j] = ans[i-2][j-1] + ans[i-2][j]
			}
		}
	}
	return ans
}
