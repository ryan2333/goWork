package main

import (
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var cmd string
	var name string
	var id int
	var line string
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Println("list")
		case "add":
			fmt.Println("add")
		}
	}

}
