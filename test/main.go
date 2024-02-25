package main

import (
	"fmt"
	"os"

	"piscine"
)

func main() {
	args := os.Args
	err := piscine.VirifyArgs(args)
	if !err {
		piscine.PrintBoard(args)
	} else {
		fmt.Println("Error")
	}
}
