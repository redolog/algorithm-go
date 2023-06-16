package app

// https://leetcode.cn/problems/number-of-times-binary-string-is-prefix-aligned/
// 1375. Number of Times Binary String Is Prefix-Aligned
// A binary string is prefix-aligned if, after the ith step, all the bits in the inclusive range [1, i] are ones and all the other bits are zeros.

// 鸽笼定理：逐个遍历，当 i+1 == max 时，说明前序[1,i+1]区间的值都已经被翻过一次了，因此可计数+1
func numTimesAllBlue(flips []int) int {
	max, cnt := flips[0], 0
	for i, flip := range flips {
		// 下标+1即为 [1,n]的每个元素值
		ele := i + 1
		if flip > max {
			max = flip
		}
		if max == ele {
			cnt++
		}
	}
	return cnt
}
