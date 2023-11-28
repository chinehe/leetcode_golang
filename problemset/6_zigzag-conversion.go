package problemset

import (
	"bytes"
)

func convert(s string, numRows int) string {
	// 特殊情况排除
	if numRows == 1 {
		return s
	}
	// 先确定第一列
	length := len(s)
	firstIndexes := make([]int, 0)
	for i := 0; i < length; i += (numRows - 1) * 2 {
		firstIndexes = append(firstIndexes, i)
	}
	// 开始拼装
	buffer := bytes.Buffer{}
	for i := 0; i < numRows; i++ {
		for _, firstIndex := range firstIndexes {
			curIndex := firstIndex + i
			if curIndex >= length {
				break
			}
			// 写入与第一行对应的行
			buffer.WriteByte(s[curIndex])
			// 如果是第一行或者最后一行
			if i == 0 || i == numRows-1 {
				continue
			}
			index := curIndex + (numRows-1-i)*2
			if index < length {
				buffer.WriteByte(s[index])
			}
		}
	}
	return buffer.String()
}
