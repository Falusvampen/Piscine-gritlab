package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := range args {
		result += args[i]
		if i != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
