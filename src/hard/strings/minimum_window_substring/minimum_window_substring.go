package minimum_window_substring

import (
	"math"
)

// https://leetcode.com/problems/minimum-window-substring/

// minWindow slid window approach [time complexity:O(n),auxiliary space:O(1)]
func minWindow(s string, t string) string {
	window, required := make(map[byte]int), make(map[byte]int)

	// init required window
	for i := range t {
		required[t[i]]++
	}

	have, need := 0, len(t)
	l, r := 0, 0
	minLen := math.MaxInt64

	for r = range s {

		// update window
		c := s[r]
		window[c]++

		// check c if required
		if window[c] == required[c] {
			have++
		}

		for have == need {

			// update minLen
			length := r - l + 1
			if length < minLen {
				minLen = length
			}

			// update window
			window[s[l]] -= 1

			// break the loop
			if window[s[l]] < required[s[l]] {
				have -= 1
			}

			l++
		}
	}

	if minLen == math.MaxInt64 {
		return ""
	} else {
		return s[l-1 : r+1]
	}
}
