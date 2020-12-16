package main

func PrintChristmasTree(numberOfLines int) (christmasTree string) {
	for line := 1; line <= numberOfLines; line++ {
		spaces := ""
		for space := 1; space <= numberOfLines-line; space++ {
			spaces += " "
		}
		stars := ""
		for star := 1; star <= line*2-1; star++ {
			stars += "*"
		}
		christmasTree += spaces + stars + spaces
		if line < numberOfLines {
			christmasTree += "\n"
		}
	}
	return
}
