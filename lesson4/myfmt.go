package main

import (
	//	"bufio"
	"fmt"
	// "os"
)

func main() {
	//输出
	s := "hello world"
	fmt.Println(s)                //打印并换行
	fmt.Print(s)                  //打印不换行
	fmt.Printf("%v", s)           //格式化输出
	fmt.Fprint("a.txt", s)        //打印并写入文件，不换行
	fmt.Fprintln("a.txt", s)      //打印并写入文件，换行
	fmt.Fprintf("a.txt", "%v", s) //格式化打印并写入文件，不换行

	//输入
	var n int           //声明变量n
	fmt.Scanf("%d", &n) //接收终输入的值，赋值给n

	var s string
	var n int
	var d int
	var line string
	f := bufio.NewReader(os.Stdin) //读取标准输入
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n') //获取整行
		fmt.Sscan(line, &s, &n, &d)  //以空格分隔，将值分别赋给指针
		if s == "stop" {
			break
		}
		fmt.Println(s, n, d) //打印后值
	}

	//不获取整行
	var cmdline, name string
	var age int
	for {
		fmt.Print("> ")
		fmt.Scan(&cmdline, &name, &age)
		fmt.Println(cmdline, name, age)
		if cmdline == "stop" {
			break
		}
	}
}
