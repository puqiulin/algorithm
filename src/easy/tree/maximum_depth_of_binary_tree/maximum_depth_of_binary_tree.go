package implement_str_str

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/

// maxDepth  approach [time complexity:O(n*m),auxiliary space:O(1)]
func maxDepth(haystack string) int {
	if haystack == needle {
		return 0
	}

	n := len(haystack)
	m := len(needle)
	if n < m {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == m-1 {
				return i
			}
		}
	}

	return -1
}

// maxDepth2  approach [time complexity:O(n+m),auxiliary space:O(m)]
func maxDepth2(haystack string) int {
	return 0
}
