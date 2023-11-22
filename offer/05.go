package offer

import (
	"regexp"
	"strings"
)

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
// 输入：s = "We are happy."
// 输出："We%20are%20happy."
// 0 <= s 的长度 <= 10000

const spaceRep = `%20`

// replaceSpaceWithForce
// 0ms 3.2MB
func replaceSpaceWithForce(s string) string {
	res := ""
	for _, ss := range s {
		if ss == ' ' {
			res += spaceRep
		} else {
			res += string(ss)
		}
	}
	return res
}

// replaceSpaceWithStringBuilder use strings.Builder
// 0ms 8MB
func replaceSpaceWithStringBuilder(s string) string {
	res := strings.Builder{}
	for _, s := range s {
		if s == ' ' {
			res.WriteString(spaceRep)
		} else {
			res.WriteRune(s)
		}
	}
	return res.String()
}

// replaceSpaceWithRegexp
// 0ms 2.1MB
func replaceSpaceWithRegexp(s string) string {
	reg := regexp.MustCompile(" ")
	return reg.ReplaceAllString(s, spaceRep)
}
