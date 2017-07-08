package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

//命名返回值，返回时可以不用指定返回的变量名
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}
func add(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(add(1, 2, 3, 4, 5))
	s := []int{1, 2, 3}
	fmt.Println(add(s...)) //...表示将切片的参数分解后传入函数，不加...会报错
}
