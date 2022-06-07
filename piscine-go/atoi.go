package piscine

func Atoi(number string) int {
	result := 0
	neg := false
	if len(number) == 0 {
		return 0
	}
	if number[0] == '+' {
		number = number[1:]
	} else if number[0] == '-' {
		neg = true
		number = number[1:]
	}
	for _, e := range number {
		if e >= '0' && e <= '9' {
			result *= 10
			result += int(e) - 48
		} else {
			return 0
		}
	}
	if neg {
		result *= -1
	}
	return result
}
