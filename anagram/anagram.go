package anagram

import (
	"maps"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	subjectMap := letterFreq(subject)
	for _, word := range candidates {
		if strings.EqualFold(subject, word) {
			continue
		}
		wordMap := letterFreq(word)
		if maps.Equal(subjectMap, wordMap) {
			result = append(result, word)
		}
	}
	return result
}

func letterFreq(word string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range strings.ToLower(word) {
		count[c]++
	}
	return count
}
