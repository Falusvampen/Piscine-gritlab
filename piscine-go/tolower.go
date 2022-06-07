package piscine

func ToLower(s string) string {
	var change rune
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			change = rune(s[i] + 32)
			runes[i] = change
		}
	}
	return string(runes)

	/*var change rune
		runes := []rune(s)
		for i := 0; i < len(s); i++ {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				change = rune(s[i] + 32)
				runes[i] = change
			}
		}
		return string(runes)
	}
	*/
}
