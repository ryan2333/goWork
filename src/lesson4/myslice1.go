//切片原理，切片相当于指针，指针的值指向原数组的内存地址所存放的值，修改指针的值，原内存地址存放的值也会发生改变
package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"a",
		"b",
		"c",
		"d",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)

}
