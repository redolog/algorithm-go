package preSum

// https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/
// 1170. Compare Strings by Frequency of the Smallest Character

// Constraints:
//
//1 <= queries.length <= 2000
//1 <= words.length <= 2000
//1 <= queries[i].length, words[i].length <= 10
//queries[i][j], words[i][j] consist of lowercase English letters.

func numSmallerByFrequency(queries []string, words []string) []int {
	var f func(s string) int
	f = func(s string) int {
		// 维护最小字典序英文字母
		minC := 'z'
		// 最小字典序字符计数
		minCnt := 0
		for _, c := range s {
			if c == minC {
				minCnt++
			} else if c < minC {
				minC = c
				minCnt = 1
			}
		}
		return minCnt
	}

	// words求f(word)，同时桶排序分桶，由于 1 <= queries[i].length, words[i].length <= 10，因此计数桶仅需10个
	wordCntBucket := make([]int, 11)
	for _, word := range words {
		wordCntBucket[f(word)]++
	}
	// 下一步尝试求每个query字符串对应f(query)，此时可定位到 wordCntBucket[f(query)]，
	// 题目求 f(queries[i]) < f(W) 计数，可转换为 wordCntBucket[>f(query)]，即 wordCntBucket中，下标在 f(query)+1（包含）之后的桶计数和
	// 此时，我们需维护一个 wordCntBucket[>f(query)] 对应的后缀和数组，12为 f(query)==10 时，求f(query)+1（包含）之后的下标哨兵
	wordCntSuffixSum := make([]int, 12)
	for i := 10; i > 0; i-- {
		wordCntSuffixSum[i] = wordCntSuffixSum[i+1] + wordCntBucket[i]
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = wordCntSuffixSum[f(query)+1]
	}
	return ans
}
