package longest_common_prefix

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/887/

// longestCommonPrefix  approach [time complexity:O(n),auxiliary space:O(1)]
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	m := len(strs[0])
	res := ""

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if len(strs[j]) == i || strs[j][i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}

	return res
}
