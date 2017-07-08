package main

import (
	"fmt"
	"io"
	"os"
)

type LineCounter struct {
	Sum int
}

func (l *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			l.Sum++
		}
	}
	return l.Sum, nil
}

func main() {
	b := new(LineCounter)
	io.Copy(b, os.Stdin)
	fmt.Println(b.Sum)
}
