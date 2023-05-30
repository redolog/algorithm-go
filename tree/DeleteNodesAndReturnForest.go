package tree

// https://leetcode.cn/problems/delete-nodes-and-return-forest/
// 1110. Delete Nodes And Return Forest

func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
	toDeleteMap := make(map[int]bool)
	for _, v := range to_delete {
		toDeleteMap[v] = true
	}

	var dfs func(node *TreeNode) *TreeNode

	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		if !toDeleteMap[node.Val] {
			// 当前节点存活
			return node
		}
		// 当前节点活不了，靠左右子孙了
		if node.Left != nil {
			ans = append(ans, node.Left)
		}
		if node.Right != nil {
			ans = append(ans, node.Right)
		}
		// 当前节点+左右子节点都检查完毕
		return nil
	}

	rootRet := dfs(root)
	if rootRet != nil {
		ans = append(ans, rootRet)
	}
	return ans
}
