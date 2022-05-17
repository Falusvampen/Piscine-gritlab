package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	// for loop only adds 1 to the count variable if there is a character from the uppercase alphabet or lowercase alphabet
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' || runes[i] >= 'a' && runes[i] <= 'z' {
			count++
		}
	}
	return count
}
