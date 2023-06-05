package listNode

func reverseListWithTraverse(head *ListNode) *ListNode {
	var pre, next *ListNode
	var curr = head

	for curr != nil {
		next = curr.Next
		curr.Next = pre
		pre = curr

		curr = next
	}
	return pre
}
func reverseListWithRecurse(curr *ListNode) *ListNode {
	if curr == nil || curr.Next == nil {
		// 返回tail尾部节点
		return curr
	}
	tailNode := reverseListWithRecurse(curr.Next)
	// 处理当前节点
	curr.Next.Next = curr
	// 原始head节点，反转完后的next为nil
	curr.Next = nil
	return tailNode
}
