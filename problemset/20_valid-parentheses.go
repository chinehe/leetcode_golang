package problemset

func isValid(s string) bool {
	stack := newStack(len(s))
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack.push(v)
		case ')', ']', '}':
			pop, err := stack.pop()
			if err != nil || pop != m[v] {
				return false
			}
		}
	}
	return stack.size() == 0
}
