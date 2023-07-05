package decimal

// https://leetcode.cn/problems/k-items-with-the-maximum-sum/
// 2600. K Items With the Maximum Sum

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if numOnes >= k {
		return k
	}
	k -= numOnes
	if numZeros >= k {
		return numOnes
	}
	k -= numZeros
	return numOnes - k
}
