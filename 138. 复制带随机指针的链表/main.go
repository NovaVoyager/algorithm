package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//https://leetcode.cn/problems/copy-list-with-random-pointer/
func main() {
	n := getNode()
	fmt.Println(copyRandomList(n))
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//生成后继节点
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}

	//连接后继节点random
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	//连接链表
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = node.Next.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}

	return headNew
}

func getNode() *Node {
	n1 := Node{
		Val:    7,
		Next:   nil,
		Random: nil,
	}
	n2 := Node{
		Val:    13,
		Next:   nil,
		Random: nil,
	}
	n3 := Node{
		Val:    11,
		Next:   nil,
		Random: nil,
	}
	n4 := Node{
		Val:    10,
		Next:   nil,
		Random: nil,
	}
	n5 := Node{
		Val:    1,
		Next:   nil,
		Random: nil,
	}

	n1.Next = &n2
	n2.Next = &n3
	n2.Random = &n1
	n3.Next = &n4
	n3.Random = &n5
	n4.Next = &n5
	n4.Random = &n3
	n5.Random = &n1

	return &n1
}
