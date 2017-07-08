//切片字面量，方法
package main

import (
	"fmt"
)

func main() {
	//切片初始化操作：先声明匿名数组，再对数组全部进行切片
	q := []int{2, 3, 5, 7, 11, 13} //切片的字面量，即切片的一种初始化方法
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	q = q[1:4]
	fmt.Println(q) //3,5,7

	q = q[:2]
	fmt.Println(q) //3,5

	q = q[1:]
	fmt.Println(q) //5

	//对原数组切片完成以后，原数组改变
	//cap表示潜在的容量，可对切片扩容，扩容不能超过原切片长度，且不需要重新分配内存，如果超过原切片潜在容量，需要重新分配内存
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) //len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]     //表示从0开始切，切到下标为0，表示空数组；潜在容量为6，没有变化
	printSlice(s) //len=0 cap=6 []

	s = s[:4]     //表示从0开始切，切到下标为4，表示前4个元素，潜在容量是6，没有变化
	printSlice(s) //len=4 cap=6 [2 3 5 7]

	s = s[2:]     //表示 从2开始切，切到切片末尾，潜在容量从2开始，潜在容量为4
	printSlice(s) //len=2 cap=4 [5 7]

	var ss []int //申明一个空切片
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil")
	}
	ss = s[:0]
	fmt.Println(ss) //[]

	//make创建切片
	a := make([]int, 5) //初始化一个int类型切片，长度为5，容量为5
	fmt.Println("a", a)

	b := make([]int, 0, 5) //初始化一个int类型切片，初始长度0，容量为5
	fmt.Println("b", b)

	c := b[:2]
	fmt.Println("c", c)

	d := c[2:5]
	fmt.Println("d", d)
	/*
		a等价于 []int{0,0,0,0,0,0}
		make([]int, 5)
	*/

	//切片append
	var sss []int
	printSlice(sss) //len=0 cap=0 []

	sss = append(sss, 0)
	printSlice(sss) //len=1 cap=1 [0]

	sss = append(sss, 1)
	printSlice(sss) //len=2 cap=2 [0 1]

	sss = append(sss, 2, 3, 4) //len=5 cap=6 [0 1 2 3 4]
	printSlice(sss)

	s1 := make([]int, 0, 1)
	_ = append(s1, 1)
	fmt.Println(s1) //[]  数组append需要用原切片接收，不然结果无效
	_ = append(s1, 2)
	fmt.Println(s1) //[]

	s2 := []int{7, 8, 9}
	sss = append(sss, s2...) //切片相加
	fmt.Println(sss)         //[0 1 2 3 4 7 8 9]

	//切片遍历
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v) //遍历下标和值
	}

	pow1 := make([]int, 10) //遍历切片下标
	for i := range pow {
		pow[i] = 1 << uint(i) //== 2**i
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value) //遍历切片值
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
