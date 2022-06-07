package piscine

// hallå, all
func Index(s string, toFind string) int {
	nice := StrLen(s)
	hello := StrLen(toFind)
	for i := 0; i <= nice-hello; i++ {
		if toFind == s[i:i+hello] {
			return i
		}
	}
	return -1
}

/*count := 0
// for loop only adds 1 to the count variable if there is a character from the uppercase alphabet or lowerca
if 0 < 1 {
	for u := range runes2 {
		for i := range runes {
			if runes[i] >= 'A' && runes[i] <= 'Z' || runes[i] >= 'a' && runes[i] <= 'z' {
				count++
			}
		}
	}
}
return count
*/
/*runes := []rune(s)
	runes2 := []rune(toFind)
	count := 0
	for i := range runes {
		for u := range runes2 {
			if runes[i] >= runes2[u] && runes[i] <= runes2[u] {
				count++
			}
		}
	}
	return count
}
*/
