package piscine

func Split(str, charset string) []string {
	var line1 int
	var line2 int
	for i := range charset {
		line1 = i + 1
	}
	for i := range str {
		line2 = i + 1
	}
	for i := 0; i < line2-line1; i++ {
		if str[i:i+line1] == charset {
			str = str[:i] + " " + str[i+line1:]
			line2 -= line1
		}
	}
	return SplitWhiteSpaces(str)
}
