package valid_sudoku

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/769/

// twoSum brute approach [time complexity:O(n*n),auxiliary space:O(1)]
func isValidSudoku(board [][]string) bool {
	notInRow := func(board [][]string, row int) bool {
		m := make(map[string]bool)

		for i := range 9 {
			if _, ok := m[board[row][i]]; ok {
				return false
			}
			if board[row][i] != "." {
				m[board[row][i]] = true
			}
		}

		return true
	}

	notInCol := func(board [][]string, col int) bool {
		m := make(map[string]bool)

		for i := range 9 {
			if _, ok := m[board[i][col]]; ok {
				return false
			}
			if board[i][col] != "." {
				m[board[i][col]] = true
			}
		}

		return true
	}

	notInBox := func(board [][]string, startRow int, startCol int) bool {
		m := make(map[string]bool)

		for i := range 3 {
			for j := range 3 {
				if _, ok := m[board[startRow+i][startCol+j]]; ok {
					return false
				}
				if board[startRow+i][startCol+j] != "." {
					m[board[startRow+i][startCol+j]] = true
				}
			}
		}

		return true
	}

	for i := range 9 {
		for j := range 9 {
			if !(notInRow(board, i) && notInCol(board, j) && notInBox(board, i-i%3, j-j%3)) {
				return false
			}
		}
	}

	return true
}
