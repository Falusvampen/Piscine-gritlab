package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, max-min)
	for i := range result {
		result[i] = min + i
	}
	return result
}
