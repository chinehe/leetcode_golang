package offer

import (
	"regexp"
)

// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
//
//数值（按顺序）可以分成以下几个部分：
//
//1.若干空格
//2.一个小数或者整数
//3.（可选）一个'e'或'E'，后面跟着一个整数
//4.若干空格

//小数（按顺序）可以分成以下几个部分：
//
//1.（可选）一个符号字符（'+' 或 '-'）
//2.下述格式之一：
//		1.至少一位数字，后面跟着一个点 '.'
//		2.至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
//		3.一个点 '.' ，后面跟着至少一位数字

//整数（按顺序）可以分成以下几个部分：
//
//1.（可选）一个符号字符（'+' 或 '-'）
//2. 至少一位数字

// 提示：
//
//1 <= s.length <= 20
//s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。

// isNumberWithRegexp use regexp
// 6.61% 6.41%
func isNumberWithRegexp(s string) bool {
	return regexp.MustCompile(`^ *(((\+|\-)?((\d+\.?\d*)|(\.\d+)))|((\+|\-)?\d+))(e((\+|\-)?\d+)|E((\+|\-)?\d+))? *$`).MatchString(s)
}

const (
	StartBlank = iota
	Sign
	Number
	NonStartPoint
	WithStartPoint
	FloatNumber
	Power
	ESign
	ENumber
	EndBlank
)

type NumberState struct {
	State int
}

func (n *NumberState) Blank() *NumberState {
	switch n.State {
	case StartBlank, EndBlank:
		return n
	case Number, WithStartPoint, FloatNumber, ENumber:
		n.State = EndBlank
		return n
	}
	return nil
}

func (n *NumberState) Number() *NumberState {
	switch n.State {
	case StartBlank, Sign, Number:
		n.State = Number
		return n
	case NonStartPoint, WithStartPoint, FloatNumber:
		n.State = FloatNumber
		return n
	case Power, ESign, ENumber:
		n.State = ENumber
		return n
	}
	return nil
}

func (n *NumberState) Sign() *NumberState {
	switch n.State {
	case StartBlank:
		n.State = Sign
		return n
	case Power:
		n.State = ESign
		return n
	}
	return nil
}

func (n *NumberState) Power() *NumberState {
	switch n.State {
	case Number, WithStartPoint, FloatNumber:
		n.State = Power
		return n
	}
	return nil
}

func (n *NumberState) Point() *NumberState {
	switch n.State {
	case StartBlank, Sign:
		n.State = NonStartPoint
		return n
	case Number:
		n.State = WithStartPoint
		return n
	}
	return nil
}

func (n *NumberState) IsNumber() bool {
	switch n.State {
	case Number, WithStartPoint, FloatNumber, ENumber, EndBlank:
		return true
	}
	return false
}

//
func isNumberWithStatusMachine(s string) bool {
	numberState := new(NumberState)
	for _, ss := range s {
		switch ss {
		case ' ':
			numberState = numberState.Blank()
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			numberState = numberState.Number()
		case 'e', 'E':
			numberState = numberState.Power()
		case '.':
			numberState = numberState.Point()
		case '+', '-':
			numberState = numberState.Sign()
		default:
			return false
		}
		if numberState == nil {
			return false
		}
	}
	return numberState.IsNumber()
}
