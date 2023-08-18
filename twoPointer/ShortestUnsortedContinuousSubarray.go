package twoPointer

import "fmt"

// https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	min, max := nums[n-1], nums[0]
	// 先找左边界 l
	l := n - 1
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			l = i
		}
	}

	// 再找右边界 r
	r := 0
	for i := 0; i < n; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			r = i
		}
	}
	fmt.Printf("l:%d r:%d", l, r)
	if l >= r {
		return 0
	}
	return r - l + 1
}
