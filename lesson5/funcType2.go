package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(m, n int) int {
	return m + n
}

func sub(m, n int) int {
	return m - n
}
func chen(m, n int) int {
	return m * n
}
func chu(m, n int) int {
	if n == 0 {
		fmt.Println("被除数不能为0")
		os.Exit(2)
	}
	return m / n
}

func main() {
	funcmap := map[string]func(int, int) int{ //操作符为key，操作符对应的函数为值
		"+": add,
		"-": sub,
		"*": chen,
		"/": chu,
	}
	m, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[3])
	f := funcmap[os.Args[2]]
	if f != nil {
		fmt.Println(f(m, n))
	}
}
