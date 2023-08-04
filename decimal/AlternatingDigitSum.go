package decimal

import "strconv"

// https://leetcode.cn/problems/alternating-digit-sum/
// 2544. Alternating Digit Sum

func alternateDigitSum(n int) int {
	nStr := strconv.Itoa(n)
	sum, sign := 0, 1
	for i := range nStr {
		sum += sign * int(nStr[i]-'0')
		sign = -sign
	}
	return sum
}
