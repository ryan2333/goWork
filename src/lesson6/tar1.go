package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	fdr, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fdr)
}
