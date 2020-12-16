package main

func PrintTannenbaum(numberOfLines int) (tannenbaum string) {

	for i := 1; i <= numberOfLines; i++ {
		// numberOfLines - i
		leerzeichen := ""
		for k := 1; k <= numberOfLines-i; k++ {
			leerzeichen += " "
		}
		// Zeile (i)
		sternchen := ""
		// Sternchen i*2-1
		for j := 1; j <= i*2-1; j++ {
			sternchen += "*"
		}
		tannenbaum += leerzeichen + sternchen + leerzeichen
		if i < numberOfLines {
			tannenbaum += "\n"
		}
	}
	return
}
