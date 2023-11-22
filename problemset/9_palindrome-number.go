package problemset

import (
	"fmt"
)

//  isPalindrome 反转数字方式
func isPalindrome(x int) bool {
	// 特殊情况排除
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return x == 0
	}
	// 开始反转数字
	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return reverted == x || reverted/10 == x
}

// isPalindrome2 转字符串方式
func isPalindrome2(x int) bool {
	// 负数直接失败
	if x < 0 {
		return false
	}
	// 转字符串比较
	xStr := fmt.Sprintf("%d", x)
	maxIndex := len(xStr) - 1
	midIndex := len(xStr) / 2
	for i := 0; i < midIndex; i++ {
		if xStr[i] != xStr[maxIndex-i] {
			return false
		}
	}
	return true
}
