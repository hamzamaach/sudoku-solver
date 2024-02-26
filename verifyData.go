package piscine

func VerifyRelatedNumbers(grid [][]int, row, column, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][column] == num {
			return false
		}
	}
	beginRow, beginColumn := (row/3)*3, (column/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[beginRow+i][beginColumn+j] == num {
				return false
			}
		}
	}
	return true
}
