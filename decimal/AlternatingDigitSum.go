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
func alternateDigitSum2(n int) int {
	sum, sign := 0, 1
	for n > 0 {
		sum += sign * (n % 10)
		n /= 10
		sign = -sign
	}
	// 观察用例：
	// n的位数为奇数时，从最低位开始到最高位sign是对称的
	// n的位数为偶数时，从最低位开始到最高位sign 与 从最高位到最低位sign 相反，因此最后乘以-1
	return sum * -sign
}
