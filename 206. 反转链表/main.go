package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/reverse-linked-list/
func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(reverseList2(l))
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseList(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	reverseHead := &ListNode{}
	tmp := reverseHead
	stackLen := len(stack) - 1
	for i := stackLen; i >= 0; i-- {
		tmp.Next = stack[i]
		tmp = tmp.Next
		if i == 0 {
			tmp.Next = nil
		}
	}

	return reverseHead.Next
}
