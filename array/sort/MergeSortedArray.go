package sort

// https://leetcode.cn/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	// i表示nums1合并之后的下标，从后往前填充
	// p1: nums1从后往前遍历指针
	// p2: nums2从后往前遍历指针
	i, p1, p2 := m+n-1, m-1, n-1
	for i >= 0 {
		if p2 < 0 || (p1 >= 0 && nums1[p1] >= nums2[p2]) {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
		i--
	}
}
