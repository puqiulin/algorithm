package rotate_image

import "slices"

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/770/

// rotate  [time complexity:O(n*n),auxiliary space:O(1)]
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	// It is equivalent to rotating the elements of diagonal
	slices.Reverse(matrix)

	//Swap elements on side of diagonal
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}
