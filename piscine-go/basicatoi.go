package piscine

func BasicAtoi(s string) int {
	var nice int = 0
	strang := []rune(s)
	for loop := 0; loop < len(strang); loop++ {
		if strang[loop] == 0 && loop == len(strang)-1 && nice == 0 {
		} else {
			nice = nice*10 + int(strang[loop]-'0')
		}
	}
	return nice
}
