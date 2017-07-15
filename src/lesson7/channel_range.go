package main

import "fmt"

func fibonacci(n int, c chan int) { //求斐波那契数死前10个数
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci1(n int, c chan int) { //求100以内的斐波那契数列
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
		if x > n {
			break
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	c1 := make(chan int, 10)
	go fibonacci1(100, c1)
	for i := range c1 {
		fmt.Println(i)
	}
}
