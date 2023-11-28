package problemset

// longestPalindrome 最长回文字串
func longestPalindrome(s string) string {
	// 先创建一个索引Map，K为字符，V为字符的索引数组，从大到小排列
	charIndexMap := make(map[uint8][]int)
	for i := len(s) - 1; i >= 0; i-- {
		if arr, ok := charIndexMap[s[i]]; ok {
			charIndexMap[s[i]] = append(arr, i)
		} else {
			charIndexMap[s[i]] = []int{i}
		}
	}
	// 定义一个方法，用于检查一个范围的字符串是否为回文字串
	check := func(x, y int) bool {
		for s[x] == s[y] {
			if x >= y {
				return true
			}
			x++
			y--
		}
		return false
	}
	// 开始遍历
	startIndex := 0
	endIndex := 0
	maxLen := 1
	for i := 0; i < len(s); i++ {
		if len(s)-i <= maxLen {
			break
		}
		indexes := charIndexMap[s[i]]
		for _, index := range indexes {
			if index-i+1 <= maxLen {
				break
			}
			if check(i, index) {
				startIndex = i
				endIndex = index
				maxLen = index - i
			}
		}
	}
	return s[startIndex : endIndex+1]
}
