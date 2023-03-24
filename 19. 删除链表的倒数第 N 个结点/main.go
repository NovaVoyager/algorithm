package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
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
	removeNthFromEnd3(l, 1)
}

//removeNthFromEnd3 双指针
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	first := head
	second := dummy

	//先移动快指针
	for i := 0; i < n; i++ {
		first = first.Next
	}

	//双指针一起移动
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}

//removeNthFromEnd2 栈
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	nodes := make([]*ListNode, 0)
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next

	return dummy.Next
}

//removeNthFromEnd 遍历法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	preHead := head
	linkLen := 1
	for preHead.Next != nil {
		linkLen++
		preHead = preHead.Next
	}

	preHead = head            //重置指针
	delPos := linkLen - n + 1 //计算删除位置
	tmp := &ListNode{
		Val:  -1,
		Next: head,
	}

	if delPos > linkLen { //超出链表范围
		return head
	}

	i := 1
	for preHead != nil {
		if i != delPos {
			i++
		} else {
			//删除正 数第一个时，需要同时更新指针
			if i == 1 {
				head = preHead.Next
			} else if i == linkLen { //最后一个
				tmp.Next = nil
			} else {
				tmp.Next = preHead.Next
			}
			break
		}

		preHead = preHead.Next
		tmp = tmp.Next
	}

	return head
}
