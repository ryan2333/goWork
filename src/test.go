package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("输入的参数不正确，example: downJx3Path 补丁包保存目录 补丁包名称")
		os.Exit(0)
	}
	s := os.Args[1:]
	fmt.Println(s)
}
