package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	PI = 3.1415926
	E  = 2.0
	D  = 9.8
)
const (
	A = iota //初台值为0开始
	B        //值为A+1
	C        //值为B+1
)

func main() {
	fmt.Println(PI, E, D)
	fmt.Println(A, B, C)
	var n int
	var f float32
	n = 10
	//f = n / 3  此行会报错，类型不匹配，n是int除以3，表达式的值是int,与f的类型float32不匹配，所以会报错
	f = float32(n) / 3
	fmt.Println(f, n)

	// n = f * 10 此行同样会报错，后面表达式的值是Float32类型，与n的类型int不匹配
	n = int(f * 3)
	fmt.Println(f, n)

	var n1 int64
	n1 = 1024127
	var n2 int8
	n2 = int8(n1)
	fmt.Println(n1, n2)

	var n3 int64
	n3 = 1024128 //int8取值范围(-128--127)
	var n4 int8
	n4 = int8(n3)
	fmt.Println(n3, n4)

	var s string
	s = strconv.Itoa(n) //int转换成字符串
	fmt.Println(s)
	n, err := strconv.Atoi(s) //字符串转换为int
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	n5, err := strconv.Atoi("123") //字符串转换为int,这里会报错
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n5)
	var x int64
	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	s = strconv.FormatInt(x, 10)
	fmt.Println(s)
	//世界上有10种人，一种懂二进制，一种不懂二进制  10表示二制制的2
}
