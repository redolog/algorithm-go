package string

// https://leetcode.cn/problems/remove-trailing-zeros-from-a-string/
// 2710. Remove Trailing Zeros From a String
// Ref: https://go101.org/article/string.html
func removeTrailingZeros(num string) string {
	n := len(num)
	tailIdx := n
	for i := range num {
		backIdx := n - i - 1
		c := num[backIdx]
		if c == '0' {
			tailIdx = backIdx
		} else {
			break
		}
	}
	return num[0:tailIdx]
}
