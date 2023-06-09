package string

// https://leetcode.cn/problems/minimize-string-length/
// 2716. Minimize String Length

func minimizedStringLength(s string) int {
	var cSet = make(map[int32]bool)
	for _, c := range s {
		cSet[c] = true
	}
	return len(cSet)
}
