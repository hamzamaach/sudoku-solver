package piscine

func VirifyArgs(args []string) bool {
	err := false
	if len(args) != 10 {
		err = true
	}
	for i, elem := range args {
		if i > 0 {
			if len(elem) != 9 {
				err = true
			}
			for j, char := range elem {
				for k, charToCompare := range elem {
					if j != k && char != '.' && char == charToCompare {
						err = true
						break
					}
				}
				if err {
					break
				}
			}
		}
	}
	return err
}
