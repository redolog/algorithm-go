package sort

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/buy-two-chocolates/
// 2706. Buy Two Chocolates

func buyChocoSort(prices []int, money int) int {
	sort.Ints(prices)
	if prices[0]+prices[1] > money {
		return money
	}
	return money - prices[0] - prices[1]
}
func buyChocoMinMin2(prices []int, money int) int {
	// 维护最小值、次小值
	min, min2 := math.MaxInt, math.MaxInt
	for _, price := range prices {
		if price < min {
			min2 = min
			min = price
		} else if price < min2 {
			min2 = price
		}
	}
	if min+min2 > money {
		return money
	}
	return money - min - min2
}
