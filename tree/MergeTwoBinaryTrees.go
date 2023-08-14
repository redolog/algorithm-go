package tree

// https://leetcode.cn/problems/merge-two-binary-trees/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	var currVal = 0
	var l1, r1 *TreeNode = nil, nil
	if root1 != nil {
		l1 = root1.Left
		r1 = root1.Right
		currVal += root1.Val
	}
	var l2, r2 *TreeNode = nil, nil
	if root2 != nil {
		l2 = root2.Left
		r2 = root2.Right
		currVal += root2.Val
	}
	curr := TreeNode{
		Val:   currVal,
		Left:  mergeTrees(l1, l2),
		Right: mergeTrees(r1, r2),
	}
	return &curr
}
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}
