package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if 2 <= len(args) {
		fmt.Println("Too many arguments")
	} else {
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(file))
	}
}
