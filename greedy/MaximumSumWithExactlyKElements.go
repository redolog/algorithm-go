package greedy

// https://leetcode.cn/problems/maximum-sum-with-exactly-k-elements/
// 2656. Maximum Sum With Exactly K Elements

func maximizeSum(nums []int, k int) int {
	var max = 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	// k个元素
	var ans = 0
	for k > 0 {
		ans += max
		max++
		k--
	}
	return ans
}
func maximizeSumWithFormula(nums []int, k int) int {
	var max = 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	// 等差数列求和
	return (k * (2*max + k - 1)) / 2
}
