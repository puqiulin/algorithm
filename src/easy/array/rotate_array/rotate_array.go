package rotate_array

import "fmt"

// RightRotateArray naive approach [time complexity:O(n*d),auxiliary space:O(1)]
func RightRotateArray(nums []int, k int) []int {
	n := len(nums)
	for k > 0 {
		last := nums[n-1]
		for i := range n - 1 {
			nums[n-i-1] = nums[n-i-2]
		}
		nums[0] = last
		k--
	}
	fmt.Println(nums)
	return nums
}

// RightRotateArray2 naive approach [time complexity:O(n),auxiliary space:O(n)]
func RightRotateArray2(nums []int, k int) []int {
	n := len(nums)
	temp := make([]int, n)
	for i := range n {
		temp[(i+k)%n] = nums[i]
	}
	nums = temp
	fmt.Println(nums)
	return temp
}

// RightRotateArray3 efficient approach [time complexity:O(n),auxiliary space:O(1)]
func RightRotateArray3(nums []int, k int) []int {
	n := len(nums)

	//Reverse the entire array
	l, r := 0, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	//Reverse the first k element
	l, r = 0, k-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	//Reverse the remain element
	l, r = k, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	fmt.Println(nums)
	return nums
}
