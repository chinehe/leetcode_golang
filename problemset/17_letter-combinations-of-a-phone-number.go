package problemset

import (
	"bytes"
)

var digitLetterMap = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	// 特殊情况处理
	if len(digits) == 0 {
		return []string{}
	}
	// 找出对应的字母列表
	digitLettersSlice := make([][]rune, len(digits))
	for i, v := range digits {
		digitLettersSlice[i] = digitLetterMap[v]
	}
	// 计算排列组合总数
	total := 1
	for _, v := range digitLettersSlice {
		total *= len(v)
	}
	mod := total
	tmp := total
	buf := make([]bytes.Buffer, total)
	for col := 0; col < len(digitLettersSlice); col++ {
		tmp /= len(digitLettersSlice[col])
		for row := 0; row < total; row++ {
			buf[row].WriteRune(digitLettersSlice[col][(row%mod)/tmp])
		}
		mod /= len(digitLettersSlice[col])
	}
	res := make([]string, total)
	for i, buffer := range buf {
		res[i] = buffer.String()
	}
	return res
}
