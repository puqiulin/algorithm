package merge_two_lists

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Value int
	Next  *ListNode
}

// mergeTwoLists two pointers approach [time complexity:O(n),auxiliary space:O(1)]
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// reduce code complexity,
	dummy := &ListNode{-1, nil}

	for l1 != nil && l2 != nil {
		if l1.Value > l2.Value {
			dummy.Next = l2
			l2 = l2.Next
		} else {
			dummy.Next = l1
			l1 = l1.Next
		}
		dummy = dummy.Next
	}

	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}

	return dummy.Next
}
