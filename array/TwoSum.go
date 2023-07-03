package array

import "sort"

// https://leetcode.cn/problems/two-sum/
// 1. Two Sum

// Constraints:
//
//2 <= nums.length <= 104
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//Only one valid answer exists.

// time:O(n*logn) space:O(logn)
func twoSumSortSolution(nums []int, target int) []int {
	// 先排序，然后双指针双向奔赴，查找target
	n := len(nums)
	// 下标数组
	indices := make([]int, len(nums))
	for i := range nums {
		indices[i] = i
	}
	// func传入less，即元素小于的规则
	sort.Slice(indices, func(i, j int) bool {
		return nums[indices[i]] < nums[indices[j]]
	})
	l, r := 0, n-1

	for l < r {
		// 顺序遍历 indices排序过的序列
		realL := indices[l]
		realR := indices[r]
		sum := nums[realL] + nums[realR]
		if sum == target {
			return []int{realL, realR}
		} else if sum < target {
			// 左侧小元素递增
			l++
		} else {
			// 右侧大元素递减
			r--
		}
	}
	return nil
}

func twoSumMapSolution(nums []int, target int) []int {
	num2Idx := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		diffIdx, contains := num2Idx[target-nums[i]]
		if contains && diffIdx != i {
			return []int{i, diffIdx}
		}
		num2Idx[nums[i]] = i
	}
	return nil
}
