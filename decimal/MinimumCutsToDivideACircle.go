package decimal

// https://leetcode.cn/problems/minimum-cuts-to-divide-a-circle/
// 2481. Minimum Cuts to Divide a Circle

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}
