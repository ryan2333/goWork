数组

package main

import (
	"fmt"
)

func main() {
	var a [3]int
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
}


//切片原理，切片相当于对内存地址的引用
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

