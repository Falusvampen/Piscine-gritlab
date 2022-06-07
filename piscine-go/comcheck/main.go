package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	count := 0
	for _, e := range args {
		if e == "01" || e == "galaxy" || e == "galaxy 01" {
			count++
		}
	}
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
