package valid_anagram

import (
	"slices"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/882/

// ReverseString myself approach [time complexity:O(n),auxiliary space:O(n)]
func isAnagram(s string, t string) bool {
	if len([]rune(s)) != len([]rune(t)) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range s {
		m1[v] += 1
	}
	for _, v := range t {
		m2[v] += 1
	}

	for i, v := range m1 {
		if v != m2[i] {
			return false
		}
	}

	return true
}

// ReverseString hashmap approach [time complexity:O(n),auxiliary space:O(n)]
func isAnagram2(s string, t string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}
	for _, v := range t {
		m[v] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

// ReverseString sorting approach [time complexity:O(n),auxiliary space:O(n)]
func isAnagram3(s string, t string) bool {
	s1 := []rune(s)
	t1 := []rune(t)
	slices.Sort(s1)
	slices.Sort(t1)
	return string(s1) == string(t1)
}
