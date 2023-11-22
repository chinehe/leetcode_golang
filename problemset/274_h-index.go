package problemset

import (
	"sort"
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	h := 0
	for paperIndex, citation := range citations {
		if citation < h {
			return h
		}
		h = min(citation, paperIndex+1)
	}
	return h
}
