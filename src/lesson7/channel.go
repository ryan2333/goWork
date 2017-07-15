package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将运算结果传入channel
}

func sum1(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
	}
	c <- sum //将运算结果传入channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) //声明channel
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //从channel读取数据
	fmt.Println(x, y, x+y)

	s1 := []string{"hello", "golang", "c++", "world"}
	c1 := make(chan string) //因为无法确定协程执行顺序，所以将结果扔到两个channel里，保证输出顺序
	c2 := make(chan string)
	go sum1(s1[:len(s1)/2], c1)
	go sum1(s1[len(s1)/2:], c2)
	t, z := <-c1, <-c2
	fmt.Printf("%s + %s = %s\n", t, z, t+z)
}
