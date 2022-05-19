package piscine

func IsPrintable(s string) bool {
	hello := []rune(s)
	count := 0
	for i := range hello {
		if 32 <= hello[i] && hello[i] <= 126 {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}
