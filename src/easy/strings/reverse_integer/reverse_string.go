package reverse_integer

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/

// ReverseString two pointers approach [time complexity:O(n),auxiliary space:O(1)]
func reverseString(s []string) []string {
	n := len(s)
	l := 0
	for l < n {
		s[l], s[n-1] = s[n-1], s[l]
		l += 1
		n -= 1
	}
	return s
}

// ReverseString intuition pointers approach [time complexity:O(n),auxiliary space:O(1)]
func reverseString2(s []string) []string {
	n := len(s)
	j := n - 1
	for i := 0; i < n/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}
