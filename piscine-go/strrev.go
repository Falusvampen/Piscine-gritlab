package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for i, strang := 0, len(runes)-1; i < strang; i, strang = i+1, strang-1 {
		runes[i], runes[strang] = runes[strang], runes[i]
	}
	return string(runes)
}
