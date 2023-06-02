package preSum

// https://leetcode.cn/problems/count-vowel-strings-in-ranges/
// 2559. Count Vowel Strings in Ranges

var vowelMap = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	// preSum[i] 表示 [0,i] 区间内的vowel单词数
	// preSum[0] 作为sentinel边界
	preSum := make([]int, n+1)
	for i, word := range words {
		preSum[i+1] = isVowel(word) + preSum[i]
	}

	m := len(queries)
	ans := make([]int, m)
	for i, query := range queries {
		l := query[0]
		r := query[1]
		ans[i] = preSum[r+1] - preSum[l]
	}
	return ans
}

func isVowel(word string) int {
	if vowelMap[word[0]] && vowelMap[word[len(word)-1]] {
		return 1
	}
	return 0
}
