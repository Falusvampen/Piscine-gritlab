package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	count := 0
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			count++
		}
		if count == len(runes) {
			return true
		}
	}
	return false
}
