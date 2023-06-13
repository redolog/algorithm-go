package array

// https://leetcode.cn/problems/shortest-distance-to-target-string-in-a-circular-array/
// 2515. Shortest Distance to Target String in a Circular Array

func closetTarget(words []string, target string, startIndex int) int {
	r, l, step, n := startIndex, startIndex, 0, len(words)
	for step < n {
		if words[r] == target || words[l] == target {
			return step
		}
		r = (r + 1) % n
		l = (l - 1 + n) % n
		step++
	}
	return -1
}
