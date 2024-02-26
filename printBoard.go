package piscine

import "fmt"

func PrintBoard(grid [][]int) {
	for _, row := range grid {
		for i, num := range row {
			if i == 8 {
				fmt.Printf("%v", num)
			} else {
				fmt.Printf("%v ", num)
			}
		}
		fmt.Printf("\n")
	}
}
