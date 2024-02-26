package string_to_integer

import (
	"math"
	"strings"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/884/

// myAtoi myself approach [time complexity:O(n),auxiliary space:O(1)]
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if s == "" {
		return 0
	}

	i, num, n, sign := 0, 0, len(s), 1
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < n {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && digit > math.MaxInt32%10) {
				if sign == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
			num = num*10 + digit
		} else {
			break
		}
		i++
	}

	return sign * num
}
