package main

import (
	"fmt"
)

func main() {
	//for循环1
	for i := 100; i < 110; i++ {
		fmt.Println(i)
	}

	//for循环2
	i := 10
	for i < 20 {
		fmt.Println(i)
		i++
	}
	//for循环（死循环）
	for {

	}
	//for循环, 遍历数组元素
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
