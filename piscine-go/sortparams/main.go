package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	var count int
	// for the for loop when reversing
	for range arguments {
		count++
	}
	for i := 1; i < count; i++ {
		for j := i + 1; j < count; j++ {
			// similar as strrev
			if arguments[i] > arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}
	// now print
	for i := 1; i < count; i++ {
		for _, argstr := range arguments[i] {
			z01.PrintRune(argstr)
		}
		z01.PrintRune('\n')
	}
}
