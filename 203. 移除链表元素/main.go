package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/remove-linked-list-elements/
func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(removeElements(l, 6))
}

func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{
		Val:  0,
		Next: head,
	}
	curr := head
	tmp := prev
	for curr != nil {
		next := curr.Next
		if curr.Val == val {
			tmp.Next = curr.Next
		} else {
			tmp = curr
		}
		curr = next
	}

	return prev.Next
}
