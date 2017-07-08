//函数类型
package main

import (
	"fmt"
	"os"
	"strconv"
)

func print() {
	fmt.Println("hello")
}

func print1() {
	fmt.Println("print1")
}

type fstruct struct { //定义函数类型的struct
	Func func()
}

func add(m, n int) int {
	return m + n
}

func sub(m, n int) int {
	return m - n
}

func main() {
	var f func()               //声明变量f是一个函数类型
	var flist [3]func()        //定义函数数组，元素为函数类型
	var fslice []func()        //定义函数切片
	var fmap map[string]func() // 定义函数map
	f = print
	f() //以函数语法调用变量
	f = print1
	f()
	funcmap := map[string]func(int, int) int{ //定义map，key为字符串，value为函数
		"+": add,
		"-": sub,
	}
	m, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[3])

	f := funcmap[os.Args[2]]
	if f != nil {
		fmt.Println(f(m, n))
	}
}
