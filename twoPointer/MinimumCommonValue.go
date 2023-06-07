package twoPointer

// https://leetcode.cn/problems/minimum-common-value/
// 2540. Minimum Common Value

func getCommon(nums1 []int, nums2 []int) int {
	// i遍历nums1，j遍历nums2
	i, j, n1, n2 := 0, 0, len(nums1), len(nums2)
	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return -1
}
