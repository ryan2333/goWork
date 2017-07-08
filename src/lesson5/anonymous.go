//匿名函数，指省略掉函数的名字，只保留入参和出参
package main

import (
	"fmt"
	"strings"
)

// func smap(r rune) rune {
// 	//fmt.Printf("%c\n", r-32)
// 	return r - 32
// }

func toupper(s string) string {
	//strings.Map(mapping func(), string) strings.map使用方法，类似python的map reduce
	return strings.Map(func(r rune) rune { //返回匿名函数得出的大写字母
		return r - ('a' - 'A')
	}, s)
}

func main() {
	// fmt.Println("test")
	// G := strings.Map(smap, "google")
	// fmt.Println(G)
	G := toupper("google")
	fmt.Println(G)
}
