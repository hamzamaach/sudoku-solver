package piscine

import "fmt"

func PrintBoard(args []string) {
	for i, elem := range args {
		if i > 0 {
			for j, char := range elem {
				if char == '.' {
					fmt.Printf("o")
					if j < 8 {
						fmt.Printf(" ")
					}
				} else {
					fmt.Printf(string(char))
					if j < 8 {
						fmt.Printf(" ")
					}
				}
			}
			fmt.Printf(string('\n'))
		}
	}
}
