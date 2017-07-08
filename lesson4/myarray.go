package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [3]int //数组初始化,此数组最多3个元素
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v) //遍历数组，打印下标和值
	}
	for _, v := range a {
		fmt.Printf("%d\n", v) //只遍历数组的值
	}

	for i := range a {
		fmt.Println(i) //只遍历数组的下标
	}
	var q [3]int = [3]int{1, 2, 3} //数组初始化,此数组最多3个元素
	q1 := [...]int{1, 2, 3, 4}     //数组初始化,此数组元素个数为后边定义几个就为几个
	q2 := [...]int{4: 2, 10: -1}   //数组初始化,长度为11，表达的为下标，第下标为4的元素值为2，下标为10的元素值为-1
	fmt.Println(q)
	fmt.Println(q1)
	fmt.Println(q2)
	var a2 [3]int
	a2 = q //数组q的值赋值给数组a2
	fmt.Println(a2)
	fmt.Println(q == a2)
	fmt.Println(&a2[0])
	fmt.Println(&q[0])
	fmt.Println(unsafe.Sizeof(a2))

}
