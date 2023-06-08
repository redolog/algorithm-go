package decimal

import "math"

// https://leetcode.cn/problems/construct-the-rectangle/
// 492. Construct the Rectangle

func constructRectangle(area int) []int {
	// 从给定面积的sqrt开始，逐步递减
	width := int(math.Sqrt(float64(area)))
	for ; area%width != 0; width-- {
	}
	return []int{area / width, width}
}
