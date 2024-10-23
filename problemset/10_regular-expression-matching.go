package problemset

func isMatch(s string, p string) bool {
	return isMatchRune([]rune(s), []rune(p))
}

func isMatchRune(s []rune, p []rune) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 {
		return false
	}
	var nextP rune
	if len(p) >= 2 {
		nextP = p[1]
	}
	if nextP != '*' && len(s) == 0 {
		return false
	}
	if nextP == '*' {
		switch p[0] {
		case '.':
			for i, _ := range s {
				if isMatchRune(s[i+1:], p[2:]) {
					return true
				}
			}
			return isMatchRune(s, p[2:])
		case '*':
			return false
		default:
			for i, v := range s {
				if v != p[0] {
					return isMatchRune(s[i:], p[2:])
				}
				if isMatchRune(s[i+1:], p[2:]) {
					return true
				}
			}
			return isMatchRune(s, p[2:])
		}
	} else {
		switch p[0] {
		case '.':
			return isMatchRune(s[1:], p[1:])
		case '*':
			return false
		case s[0]:
			return isMatchRune(s[1:], p[1:])
		default:
			return false
		}
	}

}
