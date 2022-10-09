package main

func surfaceArea(grid [][]int) int {
	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum := 0
			if grid[i][j] != 0 {
				sum += 2
			}
			if i-1 >= 0 && grid[i][j] > grid[i-1][j] {
				sum += grid[i][j] - grid[i-1][j]
			} else if i-1 < 0 {
				sum += grid[i][j]
			}
			if i+1 < len(grid) && grid[i][j] > grid[i+1][j] {
				sum += grid[i][j] - grid[i+1][j]
			} else if i+1 >= len(grid) {
				sum += grid[i][j]
			}
			if j-1 >= 0 && grid[i][j] > grid[i][j-1] {
				sum += grid[i][j] - grid[i][j-1]
			} else if j-1 < 0 {
				sum += grid[i][j]
			}
			if j+1 < len(grid[i]) && grid[i][j] > grid[i][j+1] {
				sum += grid[i][j] - grid[i][j+1]
			} else if j+1 >= len(grid[i]) {
				sum += grid[i][j]
			}
			ans += sum
		}
	}

	return ans
}
