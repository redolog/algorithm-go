package string

import "math/bits"

// https://leetcode.cn/problems/minimize-string-length/
// 2716. Minimize String Length
// Constraints:
//
//1 <= s.length <= 100
//s contains only lowercase English letters

func minimizedStringLength(s string) int {
	var cSet = make(map[int32]bool)
	for _, c := range s {
		cSet[c] = true
	}
	return len(cSet)
}
func minimizedStringLengthBit(s string) int {
	// uint is an unsigned integer type that is at least 32 bits in size
	// 使用一个32位无符号int存每个英文小写字母有无出现
	var cnt = uint(0)
	for _, c := range s {
		cnt |= 1 << (c - 'a')
	}
	return bits.OnesCount(cnt)
}
