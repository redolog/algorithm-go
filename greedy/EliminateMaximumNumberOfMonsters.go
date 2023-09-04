package greedy

import "sort"

// https://leetcode.cn/problems/eliminate-maximum-number-of-monsters/
// 1921. Eliminate Maximum Number of Monsters

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	// 怪物抵达时间
	arrivalTimeArr := make([]int, n)
	for i := 0; i < n; i++ {
		arrivalTimeArr[i] = (dist[i]-1)/speed[i] + 1
	}
	// 按抵达时间升序排列
	sort.Ints(arrivalTimeArr)
	ans := 0
	for i := 0; i < n; i++ {
		if i >= arrivalTimeArr[i] {
			return ans
		}
		ans++
	}
	return ans
}
