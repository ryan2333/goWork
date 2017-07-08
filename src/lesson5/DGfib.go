package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// func a(n int) int {
// 	if n <= 1 {
// 		return 2
// 	}
// 	return 2*a(n-1) + n - 1
// }

func main() {
	fmt.Println(fib(7))
	// fmt.Println(a(10))
}
