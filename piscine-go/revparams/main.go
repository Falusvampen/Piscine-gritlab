package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		PrintStr(args[i])
		z01.PrintRune(10)
	}
}

/*func main() {
	var count int
	arguments := os.Args[1:]
	for _ := range arguments {
		count++
		}

		for j := count; j > 0; j-- {
			z01.PrintRune(rune(e[j]))
		}
	}
	z01.PrintRune('\n')
}
*/
/*func main() {
	arguments := os.Args[1:]
	for _, element := range arguments {
		for j := len(element); j > 1; j-- {
			z01.PrintRune(rune(element[j]))
		}
		z01.PrintRune('\n')
	}
}
*/
func PrintStr(s string) {
	for i := 0; i < len([]rune(s)); i++ {
		z01.PrintRune(rune(s[i]))
	}
}
