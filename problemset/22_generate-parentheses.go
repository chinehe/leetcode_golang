package problemset

import (
	"bytes"
)

var parenthesis []string

func generateParenthesis(n int) []string {
	parenthesis = make([]string, 0)
	generate(bytes.Buffer{}, n, n)
	return parenthesis
}

func generate(buf bytes.Buffer, left, right int) {
	// 左括号已用完
	if left == 0 {
		for i := 0; i < right; i++ {
			buf.WriteRune(')')
		}
		parenthesis = append(parenthesis, buf.String())
		return
	}
	if left == right {
		// 此时只能左括号
		buf.WriteRune('(')
		generate(buf, left-1, right)
	} else {
		newBuffer := bytes.NewBuffer(buf.Bytes())
		// 写左括号
		buf.WriteRune('(')
		generate(buf, left-1, right)
		// 写右括号
		newBuffer.WriteRune(')')
		generate(*newBuffer, left, right-1)
	}
}
