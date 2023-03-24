package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/intersection-of-two-linked-lists/
func main() {
	l1 := ListNode{1, nil}
	l2 := ListNode{2, nil}
	l3 := ListNode{3, nil}
	l4 := ListNode{4, nil}
	l5 := ListNode{5, nil}
	l6 := ListNode{6, nil}
	l7 := ListNode{7, nil}
	l8 := ListNode{8, nil}

	l1.Next = &l2
	l3.Next = &l4
	l4.Next = &l5

	l6.Next = &l7
	l7.Next = &l8

	l2.Next = &l6
	l5.Next = &l6

	headA := &l1
	headB := &l3

	fmt.Println(getIntersectionNode(headA, headB))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
