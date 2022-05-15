package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for i := 0; i < nb+2; i++ {
		result = nb * i
	}
	return result
}
