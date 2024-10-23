package problemset

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if strings.HasPrefix(haystack[i:], needle) {
			return i
		}
	}
	return -1
}
