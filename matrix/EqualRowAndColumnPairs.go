package matrix

// https://leetcode.cn/problems/equal-row-and-column-pairs/
// 2352. Equal Row and Column Pairs

// 行列值逐个比较
func equalPairs(grid [][]int) int {
	n := len(grid)
	ans := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			// 检查r行、c列，依次的值是否相等
			if equal(r, c, n, grid) {
				ans++
			}
		}
	}
	return ans
}

func equal(r int, c int, n int, grid [][]int) bool {
	// i表示[0,n-1]依次遍历行、列的下标指针
	for i := 0; i < n; i++ {
		if grid[r][i] != grid[i][c] {
			return false
		}
	}
	return true
}
