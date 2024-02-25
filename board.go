package piscine

import "fmt"

type Board [9][9]int

func NewEmptyBoard() Board {
	var board Board
	return board
}

func PrintBoard(board Board) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if y == 8 {
				fmt.Print(board[x][y])
			} else {
				fmt.Print(board[x][y], " ")
			}
		}
		fmt.Println()
	}
}

