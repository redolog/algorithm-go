package decimal

import "math"

// https://leetcode.cn/problems/find-the-pivot-integer/
// 2485. Find the Pivot Integer

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		if arithmeticProgressionSum(1, i) == arithmeticProgressionSum(i, n) {
			return i
		}
	}
	return -1
}

// 设所求pivotInteger为x，需满足：sum[1,x]==sum[x,n]
// x*(x+1)/2==(n-x+1)*(n+x)/2
func pivotInteger2(n int) int {
	y := n * (n + 1) / 2
	x := int(math.Sqrt(float64(y)))
	if x*x == y {
		return x
	}
	return -1
}

// 等差数列求和，等差为1，区间为：[l,r]
func arithmeticProgressionSum(l, r int) int {
	return (r - l + 1) * (l + r)
}
