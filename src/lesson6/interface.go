package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "hello\n")
}
