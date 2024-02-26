package two_sum

import "fmt"

// twoSum brute approach [time complexity:O(n*n),auxiliary space:O(1)]
func twoSum(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// twoSum2 hashmap approach [time complexity:O(n),auxiliary space:O(n)]
func twoSum2(nums []int, target int) []int {
	n := len(nums)
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}
		m[nums[i]] = i
		fmt.Println(m)
	}

	return []int{}
}
