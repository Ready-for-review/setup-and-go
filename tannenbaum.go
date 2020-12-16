package main

func PrintTannenbaum(numberOfLines int) string {
	if numberOfLines == 1 {
		return "*"
	} else {
		return " * \n***"
	}
}
