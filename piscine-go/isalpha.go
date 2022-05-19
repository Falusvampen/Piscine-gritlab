package piscine

func IsAlpha(s string) bool {
	var count int
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= '0' && runes[i] <= '9' || runes[i] >= 'A' && runes[i] <= 'Z' {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false

	/*x := []rune(s)
	for z := range x {
		if 'A' > x[z] && x[z] < 'Z' && '0' > x[z] && x[z] < '9' || 'a' < x[z] && x[z] > 'z' {
			return false
		}

	}
	return true
	/*runes := []rune(s)
	count := 0
	// for loop only adds 1 to the count variable if there is a character from the uppercase alphabet or lowercase alphabet
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' || runes[i] >= 'a' && runes[i] <= 'z' || runes[i] > '0' && runes[i] < '9' {
			count++
		}
	}
	if count == len(s) || count == 0 {
		return true
	}
	return false*/
}
