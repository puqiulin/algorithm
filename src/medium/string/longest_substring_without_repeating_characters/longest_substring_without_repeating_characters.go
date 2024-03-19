package longest_substring_without_repeating_characters

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

// lengthOfLongestSubstring slid window approach [time complexity:O(n),auxiliary space:O(1)]
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	l, r := 0, 0
	maxLen := 0

	for r < len(s) {
		c := s[r]
		window[c]++
		r++
		for window[c] > 1 {
			window[s[l]]--
			l++
		}
		maxLen = max(r-l, maxLen)
	}
	return maxLen
}
