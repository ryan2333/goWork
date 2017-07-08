package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
	f.Close()
}
