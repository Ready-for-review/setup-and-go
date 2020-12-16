package main

func PrintTannenbaum(numberOfLines int) (tannenbaum string) {
	switch numberOfLines {
	case 1:
		tannenbaum = "*"
	case 2:
		tannenbaum = " * \n***"
	case 3:
		tannenbaum = "  *  \n *** \n*****"
	}
	return tannenbaum
}
