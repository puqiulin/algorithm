package reverse_string

import "math"

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/

// ReverseString intuition approach [time complexity:O(n),auxiliary space:O(1)]
func reverseInteger(x int) int {
	reverse := 0
	for x != 0 {
		last := x % 10
		x /= 10

		if reverse > math.MaxInt32/10 || reverse == math.MaxInt32/10 && last > 7 {
			return 0
		}
		if reverse < math.MinInt32/10 || reverse == math.MaxInt32/10 && last < -8 {
			return 0
		}

		reverse = reverse*10 + last
	}
	return reverse
}
