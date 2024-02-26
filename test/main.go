package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	err := piscine.VirifyArgs(os.Args)
	if err {
		fmt.Println("Error")
		return
	}

	grid := make([][]int, 9)
	for i := range grid {
		grid[i] = make([]int, 9)
		for j, char := range os.Args[i+1] {
			if char == '.' {
				grid[i][j] = 0
			} else if char >= '1' && char <= '9' {
				grid[i][j] = int(char - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if piscine.Solve(grid) {
		for _, row := range grid {
			for _, num := range row {
				fmt.Printf("%d ", num)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}
