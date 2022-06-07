package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	runes := []rune(str)
	nice := ""
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' || runes[i] == '\n' || runes[i] == '\t' {
			if runes[i+1] == ' ' || runes[i+1] == '\n' || runes[i+1] == '\t' {
			} else {
				result = append(result, nice)
				nice = ""
			}
		} else {
			nice = nice + string(runes[i])
		}
		if i+1 == len(runes) {
			result = append(result, nice)
		}
	}
	return result
}
