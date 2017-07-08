package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	n2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	switch os.Args[2] {
	case "+":
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("除数不能为0")
			return
		}
		fmt.Printf("%d / %d = %d\n", float32(n1), n2, n1/n2)
	}
	s := "hello"
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	for i, arg := range s {
		fmt.Printf("%d %c\n", i, arg)
	}
	s1 := "Abc中国"
	for i, arg := range s1 {
		fmt.Printf("%d %c\n", i, arg)
	}
}
