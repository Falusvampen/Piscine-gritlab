package piscine

func SortWordArr(array []string) {
	var result int
	for i := range array {
		result = i + 1
	}

	swap := 0
	for i := 0; i < result-1; i++ {
		swap = i
		for j := i + 1; j < result; j++ {
			if array[j] < array[swap] {
				swap = j
			}
		}
		array[i], array[swap] = array[swap], array[i]
	}
}
