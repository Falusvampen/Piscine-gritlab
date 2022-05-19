package piscine

func ToUpper(s string) string {
	var change rune
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			change = rune(s[i] - 32)
			runes[i] = change
		}
	}
	return string(runes)
}
