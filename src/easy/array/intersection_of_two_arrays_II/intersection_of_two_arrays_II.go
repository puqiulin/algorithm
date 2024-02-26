package intersect

import "slices"

// Intersect hashmap approach [time complexity:O(n),auxiliary space:O(n)]
func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		return Intersect(nums2, nums1)
	}

	n1, n2 := len(nums1), len(nums2)
	var temp []int
	m := make(map[int]int)

	for i := 0; i < n1; i++ {
		m[nums1[i]] += 1
	}

	for i := 0; i < n2; i++ {
		if m[nums2[i]] > 0 {
			temp = append(temp, nums2[i])
			m[nums2[i]] -= 1
		}
	}

	return temp
}

// Intersect2 two pointers approach [time complexity:O(nlogn),auxiliary space:O(n)]
func Intersect2(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	n1, n2 := len(nums1), len(nums2)
	i, j := 0, 0
	var temp []int

	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			temp = append(temp, nums1[i])
			i++
			j++
		}
	}

	return temp
}
