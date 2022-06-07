package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Scanln()
	} else {
		gotError := false
		for i, v := range args {
			if i > 1 {
				file, err := os.Open(v)
				if hasError(err) {
					gotError = true
					defer os.Exit(1)
					continue
				}
				defer file.Close()
				fileInfo, err := file.Stat()
				if hasError(err) {
					gotError = true
					defer os.Exit(1)
					continue
				}
				offset := fileInfo.Size() - int64(BasicAtoi(args[1]))
				if offset < 0 {
					offset = 0
				}
				pos := make([]byte, fileInfo.Size()-offset)
				_, err = file.ReadAt(pos, offset)
				if hasError(err) {
					gotError = true
					defer os.Exit(1)
					continue
				}
				if gotError {
					fmt.Println()
				}
				fmt.Printf("==> %v <==\n", v)
				os.Stdout.Write(pos)
				if i < len(args)-1 {
					fmt.Println()
				}
			}
		}
	}
}

func hasError(e error) bool {
	if e != nil {
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			return true
		}
	}
	return false
}

func BasicAtoi(s string) int {
	var x int = 0
	stringR := []rune(s)
	for i := 0; i < len(stringR); i++ {
		if stringR[i] == 0 && i == len(stringR)-1 && x == 0 {
		} else {
			x = x*10 + int(stringR[i]-48)
		}
	}
	return x
}
