package problemset

func lengthOfLongestSubstring(s string) int {
	startIndex := 0
	charIndexMap := make(map[rune]int)
	maxLen := 0
	length := 0
	for index, char := range s {
		if i, ok := charIndexMap[char]; ok && i >= startIndex {
			length = index - startIndex
			if length > maxLen {
				maxLen = length
			}
			startIndex = i + 1
		}
		charIndexMap[char] = index
	}
	if len(s)-startIndex > maxLen {
		return len(s) - startIndex
	}
	return maxLen
}
