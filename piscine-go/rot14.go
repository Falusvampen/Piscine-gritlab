package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'L') || (runes[i] >= 'a' && runes[i] <= 'l') {
			runes[i] = runes[i] + 14
		} else if (runes[i] >= 'M' && runes[i] <= 'Z') || (runes[i] >= 'm' && runes[i] <= 'z') {
			runes[i] = runes[i] - 12
		}
	}
	return string(runes)
}

/* func Rot14(s string) string {

	var change rune
	var change2 rune
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'A' && runes[i] < 'M' || runes[i] >= 'a' && runes[i] < 'm'{
			change = rune(s[i] + 14)
			runes[i] = change
			if runes[i] >= 'M' && runes[i] <= 'Z' || runes [i] >= 'm' && runes[i] <= 'z'{
				change = rune(s[i]-12)
				runes[]
			}
		}
		}
	}
	return string(runes)
}
*/
