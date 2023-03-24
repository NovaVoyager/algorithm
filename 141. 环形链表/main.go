package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/linked-list-cycle/
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
	l5.Next = l2

	fmt.Println("有环:", hasCycle(l1))

	//l5.Next = nil
	//fmt.Println("无环", hasCycle(l1))

}

func hasCycle(head *ListNode) bool {
	//两个指针都指向头节点
	fast, slow := head, head
	//一直循环，没有环fast会为空循环会停止，所以不会无限循环下去
	for fast != nil {
		//移动快指针
		fast = fast.Next

		//判断移动后的快指针是否还链接下一个链表
		if fast != nil {
			//再次移动
			fast = fast.Next
		}

		if fast == slow { //快指针追上慢指针，说明有环
			return true
		}

		//慢指针移动一步
		slow = slow.Next
	}

	//没有环
	return false
}
