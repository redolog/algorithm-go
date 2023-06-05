package listNode

import "fmt"

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildWithVals(vals []int) *ListNode {
	if vals == nil || len(vals) == 0 {
		//panic("ListNode.BuildWithVals can't accept empty array!!!")
		return &ListNode{Val: -1}
	}
	dummy := &ListNode{Val: -1}
	var curr, pre *ListNode = nil, dummy

	for i, val := range vals {
		curr = &ListNode{
			Val: val,
		}
		if i == 0 {
			dummy.Next = curr
		}
		pre.Next = curr
		pre = curr
	}

	return dummy.Next
}

func PrintAsStr(node *ListNode) {
	curr := node
	if curr == nil {
		return
	}
	for curr != nil {
		fmt.Print(curr.Val, "-->")
		curr = curr.Next
	}
}
