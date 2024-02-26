package contains_duplicate

import "slices"

// containsDuplicate brute force approach [time complexity:O(n*n),auxiliary space:O(1)]
func containsDuplicate(nums []int) bool {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// containsDuplicate2 sorting approach [time complexity:O(n),auxiliary space:O(1)]
func containsDuplicate2(nums []int) bool {
	n := len(nums)
	slices.Sort(nums)
	for i := range n - 1 {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// containsDuplicate3 hashset approach [time complexity:O(n),auxiliary space:O(n)]
func containsDuplicate3(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}

	return false
}
