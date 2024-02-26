package piscine

func Solve(grid [][]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if grid[row][column] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(grid, row, column, num) {
						grid[row][column] = num
						if Solve(grid) {
							return true
						}
						grid[row][column] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
