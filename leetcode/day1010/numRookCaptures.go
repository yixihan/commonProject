package main

func numRookCaptures(board [][]byte) int {
	x, y := findRook(board)

	return getAns(board, x, y)
}

func getAns(board [][]byte, x, y int) int {
	sum := 0

	for i := x; i >= 0; i-- {
		if board[i][y] == 'p' {
			sum++
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for i := x; i < len(board); i++ {
		if board[i][y] == 'p' {
			sum++
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for j := y; j >= 0; j-- {
		if board[x][j] == 'p' {
			sum++
			break
		} else if board[x][j] == 'B' {
			break
		}
	}

	for j := y; j < len(board[x]); j++ {
		if board[x][j] == 'p' {
			sum++
			break
		} else if board[x][j] == 'B' {
			break
		}
	}

	return sum
}

func findRook(board [][]byte) (x, y int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}

	return -1, -1
}
