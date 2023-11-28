package problemset

import (
	"bytes"
)

func intToRoman(num int) string {
	// 罗马矩阵
	romans := [][3]rune{
		{
			'I', 'V', 'X',
		},
		{
			'X', 'L', 'C',
		},
		{
			'C', 'D', 'M',
		},
		{
			'M', ' ',
		},
	}
	// 转换
	curIndex := 0
	runes := make([]rune, 0)
	for num > 0 {
		mod := num % 10
		big := false
		switch {
		case mod == 4:
			runes = append(runes, romans[curIndex][1])
			runes = append(runes, romans[curIndex][0])
			mod = 0
		case mod == 9:
			runes = append(runes, romans[curIndex][2])
			runes = append(runes, romans[curIndex][0])
			mod = 0
		case mod >= 5:
			big = true
			mod -= 5
		}
		for i := 0; i < mod; i++ {
			runes = append(runes, romans[curIndex][0])
		}
		if big {
			runes = append(runes, romans[curIndex][1])
		}
		num /= 10
		curIndex++
	}
	buf := bytes.Buffer{}
	for i := len(runes) - 1; i >= 0; i-- {
		buf.WriteRune(runes[i])
	}
	return buf.String()
}

func intToRoman3(num int) string {
	// 罗马矩阵
	romans := [][3]string{
		{
			"I", "V", "X",
		},
		{
			"X", "L", "C",
		},
		{
			"C", "D", "M",
		},
		{
			"M", " ",
		},
	}
	// 转换
	curIndex := 0
	res := ""
	for num > 0 {
		mod := num % 10
		switch mod {
		case 1:
			res = romans[curIndex][0] + res
		case 2:
			res = romans[curIndex][0] + romans[curIndex][0] + res
		case 3:
			res = romans[curIndex][0] + romans[curIndex][0] + romans[curIndex][0] + res
		case 4:
			res = romans[curIndex][0] + romans[curIndex][1] + res
		case 5:
			res = romans[curIndex][1] + res
		case 6:
			res = romans[curIndex][1] + romans[curIndex][0] + res
		case 7:
			res = romans[curIndex][1] + romans[curIndex][0] + romans[curIndex][0] + res
		case 8:
			res = romans[curIndex][1] + romans[curIndex][0] + romans[curIndex][0] + romans[curIndex][0] + res
		case 9:
			res = romans[curIndex][0] + romans[curIndex][2] + res
		}
		num /= 10
		curIndex++
	}
	return res
}

func intToRoman2(num int) string {
	// 罗马矩阵
	romans := [][3]string{
		{
			"I", "V", "X",
		},
		{
			"X", "L", "C",
		},
		{
			"C", "D", "M",
		},
		{
			"M", " ",
		},
	}
	// 转换
	curIndex := 0
	res := ""
	for num > 0 {
		mod := num % 10
		tmp := ""
		if mod == 9 {
			tmp = romans[curIndex][0] + romans[curIndex][2]
			mod = 0
		} else if mod == 4 {
			tmp = romans[curIndex][0] + romans[curIndex][1]
			mod = 0
		} else if mod >= 5 {
			tmp = romans[curIndex][1]
			mod -= 5
		}
		for i := 0; i < mod; i++ {
			tmp = tmp + romans[curIndex][0]
		}
		res = tmp + res
		num /= 10
		curIndex++
	}
	return res
}
