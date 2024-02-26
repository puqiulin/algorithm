package first_unique_char

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/881/

// firstUniqChar intuition approach [time complexity:O(n),auxiliary space:O(n)]
func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v] += 1
	}
	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}
