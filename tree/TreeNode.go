package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeWithIntArr(nums []int) *TreeNode {
	return BuildTreeWithIdx(nums, 0)
}

func BuildTreeWithIntArrNillable(nums []*int) *TreeNode {
	return BuildTreeWithIdxNillable(nums, 0)
}

func BuildTreeWithIdx(nums []int, i int) *TreeNode {
	if i >= len(nums) {
		return nil
	}
	var curr = &TreeNode{Val: nums[i]}
	curr.Left = BuildTreeWithIdx(nums, 2*i+1)
	curr.Right = BuildTreeWithIdx(nums, 2*i+2)
	return curr
}

func BuildTreeWithIdxNillable(nums []*int, i int) *TreeNode {
	if i >= len(nums) || nums[i] == nil {
		return nil
	}
	curr := &TreeNode{Val: *nums[i]}
	curr.Left = BuildTreeWithIdxNillable(nums, 2*i+1)
	curr.Right = BuildTreeWithIdxNillable(nums, 2*i+2)
	return curr
}

func BuildTreeWithQueue(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	// 基于链表的队列
	queue := list.New()
	root := &TreeNode{Val: nums[0]}
	queue.PushBack(root)
	i := 1
	for i < len(nums) {
		curr := queue.Front().Value.(*TreeNode)
		if i < len(nums) {
			curr.Left = &TreeNode{Val: nums[i]}
			queue.PushBack(curr.Left)
			i++
		}
		if i < len(nums) {
			curr.Right = &TreeNode{Val: nums[i]}
			queue.PushBack(curr.Right)
			i++
		}
	}
	return root
}

func BuildTreeWithQueueNillable(nums []*int) *TreeNode {
	if nums == nil || len(nums) == 0 || nums[0] == nil {
		return nil
	}
	i := 0
	// 基于链表的队列
	queue := list.New()
	root := &TreeNode{Val: *nums[i]}
	queue.PushBack(root)
	i++
	for i < len(nums) {
		curr := queue.Front().Value.(*TreeNode)
		if curr == nil {
			continue
		}
		if i < len(nums) && nums[i] != nil {
			curr.Left = &TreeNode{Val: *nums[i]}
			queue.PushBack(curr.Left)
		}
		i++
		if i < len(nums) && nums[i] != nil {
			curr.Right = &TreeNode{Val: *nums[i]}
			queue.PushBack(curr.Right)
		}
		i++
	}
	return root
}
