package piscine

func IsNumeric(s string) bool {
	nice := []rune(s)
	count := 0
	for i := range nice {
		if nice[i] <= '9' && nice[i] >= '0' {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}
