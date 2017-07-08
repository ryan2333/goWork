package main

import (
	"fmt"
)

func main() {
	var a, b int
	a = 10
	b = 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(a + 3)
	fmt.Println(a - 3)

	if a >= b {
		fmt.Println("a 大于 b")
	}
	if a > b || b > 10 {
		fmt.Println("a大于b,或者b大于10")
	}
	if (a > b && b > 3) || b > 10 {
		fmt.Println("a大于b and b大于3 或 b大于10")
	}
}
