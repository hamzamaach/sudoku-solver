package piscine

import "fmt"

func ConvertToGrid(args []string) [][]int {
	grid := make([][]int, 9)
	for i := range grid {
		grid[i] = make([]int, 9)
		for j, char := range args[i+1] {
			if char == '.' {
				grid[i][j] = 0
			} else if char >= '1' && char <= '9' {
				grid[i][j] = int(char - '0')
			} else {
				fmt.Println("Error")
			}
		}
	}
	return grid
}
