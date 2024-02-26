package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	// vérifier si les arguments donnée est valide
	err := piscine.VirifyArgs(os.Args)
	if err {
		fmt.Println("Error")
		return
	}
	// convert args (type dyalhom []string) given to grid (type dyalo [][]int)
	grid := piscine.ConvertToGrid(os.Args)

	if piscine.Solve(grid) {
		piscine.PrintBoard(grid)
	} else {
		fmt.Println("Error")
	}
}
