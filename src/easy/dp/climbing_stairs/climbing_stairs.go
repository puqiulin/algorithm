package climbing_stairs

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/569/

func climb(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	memo[n] = climb(n-1, memo) + climb(n-2, memo)
	return memo[n]
}

// climbStairs memo approach [time complexity:O(n),auxiliary space:O(1)]
func climbStairs(n int) int {
	m := make(map[int]int)
	return climb(n, m)
}
