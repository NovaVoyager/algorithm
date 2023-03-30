package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/odd-even-linked-list/
func main() {
	//l := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: &ListNode{
	//						Val: 6,
	//						Next: &ListNode{
	//							Val:  7,
	//							Next: nil,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	//l2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}

	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(oddEvenList(l3))
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	evenP, oddP := &ListNode{}, &ListNode{}
	tmp := head
	headEvenP, headOddP := evenP, oddP
	i := 1
	for tmp != nil {
		if i%2 == 0 { //偶数
			evenP.Next = tmp
			evenP = evenP.Next
		} else { //奇数
			oddP.Next = tmp
			oddP = oddP.Next
		}
		tmp = tmp.Next
		i++
	}

	evenP.Next = nil
	oddP.Next = headEvenP.Next

	return headOddP.Next
}
