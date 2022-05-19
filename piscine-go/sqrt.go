package piscine

func Sqrt(nb int) int {
	// for loop that goes to the value of the number and then looks if it is equal to the number when multiplied with itself
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	// otherwise returns 0, it will only return 0 for whole numbers.
	return 0
}
