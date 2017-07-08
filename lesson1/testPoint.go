package main

import "fmt"

func main() {
	var x int
	x = 1
	var y int
	y = 2
	swap(&x, &y)
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

func swap(p *int, q *int) {
	var t = *p
	*p = *q
	*q = t
}
