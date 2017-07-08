package main

import "fmt"

func add(x int, y int) int { //传入参数及类型，返回值类型
	return x + y
}

func f(i, j, k int, s, t string) (int, int, int, string, string) { //多个参数同类型，可以省略，多返回值需要用小括号括起来
	//等价于func f(i int ,j int, k int, s string, t string)
	return i, j, k, s, t
}

func split(sum int) (x, y int) { //命名返回值，明确指定返回哪些变量的值
	x = sum / 10
	y = sum % 10
	return
}

func add1(args ...int) int { //可变参数,...放前面表示聚合，将参数取成一个切片，可以进行遍历
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

func fib(n int) int { //递归函数，代码简洁但效率低下
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(add(42, 33))
	fmt.Println(f(1, 2, 3, "hh", "kk"))
	fmt.Println(split(19))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(add1(s...)) //...放后面表示将切片的参数分解后传入函数，不加...会报错
}
