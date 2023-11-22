package problemset

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	elemFrequentMap := make(map[int]int, 0)
	for i := range nums {
		elemFrequentMap[nums[i]]++
	}
	type pair struct {
		num      int
		frequent int
	}
	pairs := make([]pair, 0)
	for num, frequent := range elemFrequentMap {
		pairs = append(pairs, pair{num, frequent})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].frequent > pairs[j].frequent
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = pairs[i].num
	}
	return res
}
