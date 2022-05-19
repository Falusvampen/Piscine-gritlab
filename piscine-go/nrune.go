package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n > 0 && n <= len(s) {
		return runes[n-1]
	}
	return 0
}
