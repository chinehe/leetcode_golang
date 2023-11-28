package problemset

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	// 去除前后空格
	s = strings.Trim(s, " ")
	if s == "" {
		return 0
	}
	// 获取正负号
	negative := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		negative = -1
	}
	// 获取数字
	num := 0
	// 数字
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')*negative
			if num > math.MaxInt32 {
				return math.MaxInt32
			} else if num < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			// 非数字直接结束
			break
		}
	}
	return num
}
