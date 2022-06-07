package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, v := range args {
			file, err := ioutil.ReadFile(v)
			if err != nil {
				os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
				os.Exit(1)
			}
			os.Stdout.Write(file)
		}
	}
}
