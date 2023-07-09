package twoPointer

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 167. Two Sum II - Input Array Is Sorted

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			// 和小，那么我们需要l右挪，稍微变大一点
			l++
		} else {
			r--
		}
	}
	return nil
}

func twoSumBSearchSolution(numbers []int, target int) []int {
	NOT_FOUND := -1
	// 给定排序数组，通过二分的方式找到target位置，找不到时返回-1
	bSearch := func(arr []int, l, r, target int) int {
		for l <= r {
			mid := l + ((r - l) >> 1)
			if arr[mid] == target {
				return mid
			} else if arr[mid] > target {
				r--
			} else {
				l++
			}
		}
		return NOT_FOUND
	}
	n := len(numbers)
	for i, a := range numbers {
		bIdx := bSearch(numbers, i+1, n-1, target-a)
		if bIdx != NOT_FOUND {
			return []int{i + 1, bIdx + 1}
		}
	}
	return nil
}
