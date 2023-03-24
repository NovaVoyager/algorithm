package main

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/merge-two-sorted-lists/
func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists(l1, l2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil && list2 != nil:
		return list2
	case list1 != nil && list2 == nil:
		return list1
	case list1 == nil && list2 == nil:
		return nil
	}

	preHead := &ListNode{}
	tmpPre := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmpPre.Next = list1
			list1 = list1.Next
		} else {
			tmpPre.Next = list2
			list2 = list2.Next
		}
		tmpPre = tmpPre.Next
	}

	if list1 == nil {
		tmpPre.Next = list2
	} else if list2 == nil {
		tmpPre.Next = list1
	}

	return preHead.Next
}
