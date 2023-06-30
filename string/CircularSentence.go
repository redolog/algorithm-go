package string

import "strings"

// https://leetcode.cn/problems/circular-sentence/
// 2490. Circular Sentence

func isCircularSentence(sentence string) bool {
	n := len(sentence)
	for i := 0; i < n; i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return sentence[0] == sentence[n-1]
}
func isCircularSentence2(sentence string) bool {
	wordArr := strings.Split(sentence, " ")
	n := len(wordArr)
	for i := 0; i < n; i++ {
		if wordArr[i][0] != wordArr[(i-1+n)%n][len(wordArr[(i-1+n)%n])-1] {
			return false
		}
	}
	return true
}
