package main

import (
	"fmt"
)

func toupper(s string) string {
	array1 := []byte(s)
	for i := 0; i < len(array1); i++ {
		if array1[i] >= 'a' && array1[i] < 'z' {
			array1[i] -= 32
		} else if array1[i] < 'a' && array1[i] >= 'A' {
			array1[i] += 32
		}
	}
	s = string(array1)
	return s
}

func main() {
	str1 := "hello" + "world" //字符串相加
	str2 := "helloworld"
	doc := `
			即使换行也不影响
			可以验证一下
			类似python的'''
		`

	fmt.Println(str1)
	fmt.Println(doc)
	if str1 == str2 {
		fmt.Println("equal")
	}
	var c1 byte  //声明变量，byte类型
	c1 = str1[0] //切片，取str1第一个字符
	fmt.Println(str1, str2, c1)
	fmt.Printf("%d %c\n", c1, c1)

	str3 := str1[3:4] //切片
	fmt.Println(str3)

	//打印ascii表
	var b byte
	for b = 0; b < 20; b++ {
		fmt.Printf("%d %c\n", b, b)
	}

	array := []byte(str1)
	fmt.Println(array)
	array[0] = 'H' //H=72
	str1 = string(array)
	fmt.Println(str1)

	fmt.Println('a' + 'H' - 'h')
	fmt.Println(toupper("HelloWorld"))
}
