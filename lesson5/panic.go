package main

import (
	"fmt"
)

func print() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	var p *int
	fmt.Println(*p)
}

func main() {

	//	var n int
	//	fmt.Println(10 / n) //除数为0也会产生panic
	print()
	panic("don't want exec continue")
	var i = 3
	var slice [3]int
	fmt.Println(slice[i])
}
