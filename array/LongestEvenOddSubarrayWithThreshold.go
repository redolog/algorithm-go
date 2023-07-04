package array

// https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/
// 2760. Longest Even Odd Subarray With Threshold

// Find the length of the longest subarray of nums starting at index l and ending at index r (0 <= l <= r < nums.length) that satisfies the following conditions:
//
//nums[l] % 2 == 0
//For all indices i in the range [l, r - 1], nums[i] % 2 != nums[i + 1] % 2
//For all indices i in the range [l, r], nums[i] <= threshold
//Return an integer denoting the length of the longest such subarray.

// 寻找区间 [l,r]，满足：
// 1. l边界数值为偶数；
// 2. [l,r-1]区间内，任意一个元素奇偶性与相邻元素奇偶性不同；
// 3. [l,r]区间内，任意一项数值不大于 threshold；
// 4. 求这样的子数组最大的长度；
func longestAlternatingSubarray(nums []int, threshold int) int {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	ans, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i]%2 != 0 || nums[i] > threshold {
			continue
		}
		// 锁定l
		l := i
		i++
		// 找r
		for ; i < n; i++ {
			if nums[i] > threshold {
				break
			}
			if nums[i-1]%2 == nums[i]%2 {
				break
			}
		}
		ans = max(ans, i-l)
		// 此时r需从头判定，因此-1
		i--
	}
	return ans
}
