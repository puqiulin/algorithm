package move_zeroes

// moveZeroes bubble sort approach [time complexity:O(n),auxiliary space:O(1)]
func moveZeroes(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
