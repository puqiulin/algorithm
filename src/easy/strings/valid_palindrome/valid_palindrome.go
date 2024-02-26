package valid_palindrome

import (
	"strings"
	"unicode"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/883/

// isPalindrome two pointers approach [time complexity:O(n),auxiliary space:O(1)]
func isPalindrome(s string) bool {
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)

	s1 := []rune(s)
	l, r := 0, len(s1)-1

	for l < r {
		if s1[l] != s1[r] {
			return false
		}
		l++
		r--
	}

	return true
}
