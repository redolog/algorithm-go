package listNode

// https://leetcode.cn/problems/add-two-numbers-ii/
// 445. Add Two Numbers II

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 反转链表
	l1 = reverse(l1)
	l2 = reverse(l2)
	// 2. 根据两数加和 addTwoNumbers 求和，得到l3
	l3 := addTwoNumbers(l1, l2)
	// 3. l1 l2 再次反转，归位
	l1 = reverse(l1)
	l2 = reverse(l2)
	// 4. 返回 l3
	return reverse(l3)
}

func reverse(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}
	last := reverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return last
}

func addTwoNumbersIIListSolution(l1 *ListNode, l2 *ListNode) *ListNode {
	// 使用列表存储入参数据
	// 倒序加和
	// 倒序的同时构建答案
	l1Vals, l2Vals := node2List(l1), node2List(l2)
	n1, n2, carry := len(l1Vals), len(l2Vals), 0

	// 创建一个指针对象
	//curr := &ListNode{}
	var curr *ListNode
	// next类型为ListNode指针类型，且初始为nil
	var next *ListNode
	for n1 > 0 || n2 > 0 || carry > 0 {
		val := carry
		if n1 > 0 {
			val += l1Vals[n1-1]
			n1--
		}
		if n2 > 0 {
			val += l2Vals[n2-1]
			n2--
		}
		carry = val / 10
		curr = &ListNode{
			Val:  val % 10,
			Next: next,
		}
		next = curr
	}
	return curr
}

func node2List(node *ListNode) []int {
	var list []int
	for node != nil {
		list = append(list, node.Val)
		node = node.Next
	}
	return list
}
