package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, e := range arguments {
		for j := 0; j < len([]rune(e)); j++ {
			z01.PrintRune(rune(e[j]))
		}
		z01.PrintRune('\n')
	}
}
