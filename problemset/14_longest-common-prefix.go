package problemset

func longestCommonPrefix(strs []string) string {
	// 最小字符串长度
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	// 开始最长公共前缀
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLen]
}
