package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	transformed := make(map[string]int)
	for score, letters := range in {
		for _, letter := range letters {
			transformed[strings.ToLower(letter)] = score
		}
	}
	return transformed
}
