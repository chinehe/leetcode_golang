package problemset

type variableNode struct {
	x int
	y int
	v []int
}

func solveSudoku(board [][]byte) {
	// 行、列、块可选项
	rowElems := [9]map[byte]struct{}{}
	columnElems := [9]map[byte]struct{}{}
	blockElems := [3][3]map[byte]struct{}{}
	// 遍历初始化
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			rowElems[i][board[i][j]] = struct{}{}
			columnElems[j][board[i][j]] = struct{}{}
			blockElems[i/3][j/3][board[i][j]] = struct{}{}
		}
	}
}
