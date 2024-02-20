package remove_duplicates

import "fmt"

// Naive Approach (Using extra space)
func RemoveDuplicates(nums []int) int {
	// Return if array is empty or only a single element
	n := len(nums)
	if n <= 1 {
		return n
	}

	temp := make([]int, n)
	index := 0
	for i := range n - 1 {
		// Store the current element if different from next
		if nums[i] != nums[i+1] {
			temp[index] = nums[i]
			index++
		}
	}

	// Store the last one element
	temp[index] = nums[n-1]
	index++

	// Modify original array
	for i := range index {
		nums[i] = temp[i]
	}

	fmt.Printf("nums is %v\n", nums)

	return index
}

// Efficient Approach (Maintaining a separate index for the same array )
func RemoveDuplicates2(nums []int) int {
	// Return if array is empty or only a single element
	n := len(nums)
	if n <= 1 {
		return n
	}

	index := 0
	for i := range n - 1 {
		// Store the current element in same array if different from next
		if nums[i] != nums[i+1] {
			nums[index] = nums[i]
			index++
		}
	}

	// Store the last one element
	nums[index] = nums[n-1]
	index++

	fmt.Printf("nums is %v\n", nums)

	return index
}
