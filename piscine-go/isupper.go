package piscine

func IsUpper(s string) bool {
	var count int
	runes := []rune(s)
	for _, i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
	/*runes := []rune(s)
	count := 0
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			count++
		}
		if count == len(runes) {
			return true
		}
	}
	return false */
}
