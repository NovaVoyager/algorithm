package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/palindrome-linked-list/
func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	fmt.Println(isPalindrome(l))
}

func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	headIndex, endIndex := 0, len(vals)-1
	for headIndex <= endIndex {
		if vals[headIndex] != vals[endIndex] {
			return false
		}
		headIndex++
		endIndex--
	}

	return true
}
