package problemset

import (
	"bytes"
	"sort"
)

func frequencySort(s string) string {
	charNumMap := make(map[byte]int, 0)
	for i := range s {
		charNumMap[s[i]]++
	}
	type pair struct {
		char byte
		num  int
	}
	pairs := make([]pair, 0)
	for char, num := range charNumMap {
		pairs = append(pairs, pair{
			char: char,
			num:  num,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].num > pairs[j].num
	})
	builder := bytes.Buffer{}
	for i := range pairs {
		builder.Write(bytes.Repeat([]byte{pairs[i].char}, pairs[i].num))
	}
	return builder.String()
}
