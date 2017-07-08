package main

import (
	"fmt"
	"log"
	"os"
)

func print() {
	defer func() { //defer是在函数结束之前执行，作用：文件关闭，关闭锁，关闭channel,关闭socket
		fmt.Println("defer")
	}()
	fmt.Println("hello")

}

func readFile() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() //defer关闭文件
}

func main() {
	print()
}
