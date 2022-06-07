package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	var nice int
	for i := 0; i < len(s); i++ {
		nice = i
	}
	return rune(runes[nice])
}
