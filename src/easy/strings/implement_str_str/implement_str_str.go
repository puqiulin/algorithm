package implement_str_str

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/885/

// ReverseString intuitive approach [time complexity:O(n*m),auxiliary space:O(1)]
func strStr(haystack string, needle string) int {
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

// https://www.youtube.com/watch?v=JoF0Z7nVSrA
// ReverseString kmp approach [time complexity:O(n+m),auxiliary space:O(m)]
func strStr2(haystack string, needle string) int {
	return 0
}
