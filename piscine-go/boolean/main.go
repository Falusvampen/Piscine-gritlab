package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	yes := true
	no := false
	if nbr%2 == 0 {
		return yes
	} else {
		return no
	}
}

func main() {
	lengthOfArg := 0
	// a := 1
	// b := 2
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	for range os.Args[1:] {
		lengthOfArg++
	}
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
