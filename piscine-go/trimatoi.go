package piscine

func TrimAtoi(s string) int {
	minus := false
	var res int
	res = 0
	for i, v := range s {
		if res == 0 && v == '-' {
			minus = true
		}
		if v >= '0' && v <= '9' && i >= 0 {
			res = res*10 + int(v-48)
		}
	}
	if minus {
		res = res - 2*res
	}
	return res
}
