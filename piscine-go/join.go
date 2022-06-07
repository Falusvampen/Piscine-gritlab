package piscine

func Join(strs []string, sep string) string {
	strings := ""
	for i := 0; i < len(strs); i++ {
		strings = strings + strs[i]
		if i != len(strs)-1 {
			strings = strings + sep
		}
	}
	return strings
}
