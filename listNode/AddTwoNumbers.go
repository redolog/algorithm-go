package listNode

// https://leetcode.cn/problems/add-two-numbers/
// 2. Add Two Numbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := &ListNode{}
	prev := dummy
	// 进位
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {

		currVal := carry
		if l1 != nil {
			currVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currVal += l2.Val
			l2 = l2.Next
		}

		if currVal > 9 {
			carry = 1
			currVal = currVal % 10
		} else {
			carry = 0
		}
		curr = &ListNode{Val: currVal}
		prev.Next = curr
		prev = curr
		curr = curr.Next
	}
	return dummy.Next
}

func addTwoNumbersRecurseSolution(l1 *ListNode, l2 *ListNode) *ListNode {
	return recurse(l1, l2, 0)
}

func recurse(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	// 构造当前节点值并返回
	currVal := carry
	if l1 != nil {
		currVal += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		currVal += l2.Val
		l2 = l2.Next
	}
	curr := &ListNode{Val: currVal % 10}
	carry = currVal / 10
	curr.Next = recurse(nextNullable(l1), nextNullable(l2), carry)
	return curr
}

func nextNullable(l1 *ListNode) *ListNode {
	if l1 == nil {
		return nil
	}
	return l1
}
