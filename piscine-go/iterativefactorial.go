package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb >= 21 {
		return 0
	}
	if nb < -1 {
		return 0
	}
	result := nb
	for i := 1; i < nb; i++ {
		result = result * i
	}
	return result
}
