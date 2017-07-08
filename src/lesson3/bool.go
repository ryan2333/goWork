package main

import (
	"fmt"
)

func main() {
	var b bool
	b = false
	b = ("hello" == "world")

	if b {
		fmt.Println("true")
	}
}
