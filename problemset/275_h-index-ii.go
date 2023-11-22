package problemset

import (
	"sort"
)

func hIndexII(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
}
