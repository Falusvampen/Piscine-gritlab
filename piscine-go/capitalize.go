package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	if runes[0] >= 97 && runes[0] <= 122 { // om första tecknet är alphanumeriskt så höjer vi det direkt
		runes[0] -= 32
	}
	for i := range runes {
		if !((runes[i] >= 48 && runes[i] <= 57) || (runes[i] >= 65 && runes[i] <= 90) || (runes[i] >= 97 && runes[i] <= 122)) { /* vi kontrollerar om karaktären är inte alfanumeriskt, om det är nästa alfanumeriska tecken är versaler*/
			if i < len(runes)-1 {
				if runes[i+1] >= 97 && runes[i+1] <= 122 {
					runes[i+1] -= 32
				}
			}
		} else { // om tecknet är alfanumeriskt kontrollerar vi om det inte står i början av ordet, i så fall sätter vi det med gemener
			if i > 0 {
				if (runes[i-1] >= 48 && runes[i-1] <= 57) || (runes[i-1] >= 65 && runes[i-1] <= 90) || (runes[i-1] >= 97 && runes[i-1] <= 122) {
					if runes[i] >= 65 && runes[i] <= 90 {
						runes[i] += 32
					}
				}
			}
		}
	}
	return string(runes)
}
