package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	sort.Slice(args, func(i, j int) bool {
		return args[i] < args[j]
	})
	for _, res := range args {
		fmt.Println(res)
	}
}
