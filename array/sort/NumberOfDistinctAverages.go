package sort

import "sort"

// https://leetcode.cn/problems/number-of-distinct-averages/
// 2465. Number of Distinct Averages

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	sumSet := make(map[int]bool)
	for l := 0; l < n>>1; l++ {
		sumSet[nums[l]+nums[n-1-l]] = true
	}
	return len(sumSet)
}
