package piscine

func IsUpper(s string) bool {
	var count int
	runes := []rune(s)
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
