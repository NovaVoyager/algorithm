package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func main() {
	//stackTest()
	node := getNode()
	fmt.Println(flatten2(node))
}

func flatten2(root *Node) *Node {
	dfs2(root)
	return root
}

func dfs2(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		next := cur.Next
		//优先处理子节点
		if cur.Child != nil {
			childLast := dfs2(cur.Child)

			next = cur.Next

			//连接当前节点与子节点
			cur.Next = cur.Child
			cur.Child.Prev = cur

			//连接下一个节点与最后子节点
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}

			cur.Child = nil //置空子节点
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}

	return
}

func dfs(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		next := cur.Next
		// 如果有子节点，那么首先处理子节点
		if cur.Child != nil {
			childLast := dfs(cur.Child)

			next = cur.Next
			// 将 node 与 child 相连
			cur.Next = cur.Child
			cur.Child.Prev = cur

			// 如果 next 不为空，就将 last 与 next 相连
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}

			// 将 child 置为空
			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}
	return
}

func flatten(root *Node) *Node {
	curr := root
	stack := Stack{}
	for curr != nil {
		if curr.Child != nil {
			child := curr.Child
			stack.Push(curr.Next)
			curr.Next = curr.Child
			curr.Child.Prev = curr
			curr.Child = nil
			curr = child
		}

		if curr.Child != nil { //此步目的判断当前节点是否有子节点，有就继续
			continue
		}

		if curr.Next == nil && len(stack.value) > 0 {
			stackNode := stack.Pop()
			if stackNode == nil { //此步目的判断弹出的节点是否为nil，为nil就继续
				continue
			}
			curr.Next = stackNode
			stackNode.Prev = curr
		}

		curr = curr.Next
	}

	return root
}

type Stack struct {
	value []*Node
}

func (s *Stack) Push(value *Node) {
	if s.value == nil {
		s.value = make([]*Node, 0)
	}
	s.value = append(s.value, value)
}

func (s *Stack) Pop() *Node {
	if len(s.value) == 0 {
		return nil
	}
	stackHead := len(s.value) - 1
	node := s.value[stackHead]
	if stackHead == 0 {
		s.value = nil
	} else {
		s.value = s.value[:stackHead]
	}
	return node
}

func getNode2() *Node {
	l1 := Node{
		Val:   1,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l2 := Node{
		Val:   2,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l3 := Node{
		Val:   3,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	l1.Child = &l2
	l2.Child = &l3

	return &l1
}

func getNode() *Node {
	l1 := Node{
		Val:   1,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l2 := Node{
		Val:   2,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l3 := Node{
		Val:   3,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l4 := Node{
		Val:   4,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l5 := Node{
		Val:   5,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l6 := Node{
		Val:   6,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l7 := Node{
		Val:   7,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l8 := Node{
		Val:   8,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l9 := Node{
		Val:   9,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l10 := Node{
		Val:   10,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l11 := Node{
		Val:   11,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	l12 := Node{
		Val:   12,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	l1.Next = &l2
	l2.Prev = &l1
	l2.Next = &l3
	l3.Prev = &l2
	l3.Next = &l4
	l4.Prev = &l3
	l4.Next = &l5
	l5.Prev = &l4
	l5.Next = &l6
	l6.Prev = &l5

	l7.Next = &l8
	l8.Prev = &l7
	l8.Next = &l9
	l9.Prev = &l8
	l9.Next = &l10
	l10.Prev = &l9

	l11.Next = &l12
	l12.Prev = &l11

	l3.Child = &l7
	l8.Child = &l11

	return &l1
}

func stackTest() {
	stack := Stack{}

	a := &Node{
		Val:   1,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	b := &Node{
		Val:   2,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	c := &Node{
		Val:   3,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	stack.Push(a)
	stack.Push(b)
	stack.Push(c)

	stack.Pop()
	stack.Pop()
	stack.Pop()
}
