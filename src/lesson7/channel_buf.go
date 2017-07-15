package main

import "fmt"

func sum1(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
	}
	c <- sum //将运算结果传入channel
}
func main() {
	s1 := []string{"hello", "golang", "c++", "world"}
	c1 := make(chan string, 10) //声明channel添加buf,存数据，会么时候需要什么时候取
	// c2 := make(chan string)
	go sum1(s1[:len(s1)/2], c1)
	go sum1(s1[len(s1)/2:], c1)
	t, z := <-c1, <-c1
	fmt.Printf("%s + %s = %s\n", t, z, t+z)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
