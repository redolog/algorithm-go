package decimal

// https://leetcode.cn/problems/average-value-of-even-numbers-that-are-divisible-by-three/
// 2455. Average Value of Even Numbers That Are Divisible by Three
func averageValue(nums []int) int {
	sum := 0
	cnt := 0
	for _, num := range nums {
		if num%6 == 0 {
			sum += num
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
