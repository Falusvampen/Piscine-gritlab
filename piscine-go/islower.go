package piscine

func IsLower(s string) bool {
	var count int
	runes := []rune(s)
	for _, i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			count++
		}
		if count == len(runes) {
			return true
		}
	}
	return false
}

/*count := 0
runes := []rune(s)
for i := range runes {
	if runes[i] >= 'a' && runes[i] <= 'z' {
		count++
	}
	if count == len(runes) {
		return true
	}
}
return false*/
