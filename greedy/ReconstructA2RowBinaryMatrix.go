package greedy

// https://leetcode.cn/problems/reconstruct-a-2-row-binary-matrix/
// 1253. Reconstruct a 2-Row Binary Matrix

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	ans := make([][]int, 2)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	// 先往前填充1，消费upper、lower
	for i := 0; i < n; i++ {
		colSumCurr := colsum[i]
		if colSumCurr == 0 {
			continue
		}
		row0, row1 := ans[0], ans[1]
		if colSumCurr == 2 {
			row0[i] = 1
			row1[i] = 1
			upper--
			lower--
		} else {
			if upper > lower {
				// 优先消费upper
				row0[i] = 1
				upper--
			} else {
				// 优先消费lower
				row1[i] = 1
				lower--
			}
		}
	}
	if !(upper == 0 && lower == 0) {
		return [][]int{}
	}
	return ans
}
