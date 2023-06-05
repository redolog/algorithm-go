package app

// https://leetcode.cn/problems/apply-operations-to-an-array/
// 2460. Apply Operations to an Array

func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	i := 0
	ans := make([]int, n)
	for _, num := range nums {
		if num > 0 {
			ans[i] = num
			i++
		}
	}
	return ans
}

// 一边执行赋值操作，一边维护数组左侧非零元素
func applyOperations2(nums []int) []int {
	// i用于遍历完整数组：[0,n-1]
	// j用于维护数组左侧非零下标
	n := len(nums)
	for i, j := 0, 0; i < n; i++ {
		if i+1 < n && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
