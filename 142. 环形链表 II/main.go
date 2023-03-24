package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/linked-list-cycle-ii/
func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l1

	fmt.Println("有环:", detectCycle(l1))

}

func detectCycle(head *ListNode) *ListNode {
	//遍历链表，把每一个节点放入hash表
	//当hash出现重复时，说明有环，重复的链表节点也为入口节点
	tmp := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := tmp[head]; ok {
			return head
		}

		tmp[head] = struct{}{}
		head = head.Next
	}

	return nil
}
