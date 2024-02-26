package single_number

import (
	"slices"
)

// SingleNumber brute force approach [time complexity:O(n*n),auxiliary space:O(1)]
func SingleNumber(nums []int) int {
	n := len(nums)
	var flag bool

	for i := 0; i < n; i++ {
		flag = true
		for j := 0; j < n; j++ {
			if !flag {
				break
			}
			if i == j {
				continue
			}
			if nums[i] == nums[j] {
				flag = false
			}
		}
		if flag {
			return nums[i]
		}
	}

	return -1
}

// SingleNumber2 sorting approach [time complexity:O(nlogn),auxiliary space:O(1)]
func SingleNumber2(nums []int) int {
	n := len(nums)
	slices.Sort(nums)

	i := 0
	for i < n-2 {
		if nums[i] == nums[i+1] {
			i += 2
		} else {
			return nums[i]
		}
	}

	return nums[n-1]
}

// SingleNumber3 hashmap approach [time complexity:O(n),auxiliary space:O(n)]
func SingleNumber3(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}

	for i := range m {
		return i
	}

	return -1
}

// SingleNumber4 bitwise xor approach [time complexity:O(n),auxiliary space:O(1)]
func SingleNumber4(nums []int) int {
	res := 0 // n ^ 0 = n
	for _, n := range nums {
		res = n ^ res
	}
	return res
}
