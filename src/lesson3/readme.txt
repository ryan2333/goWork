数据类型(强类型)
	整数
	字符串
	布尔
	浮点

整数：
	int, int32, int64, uint, uint32, uint64   u表示unsigned(无符号整数)

布尔型：
	值只有true 和 false

浮点数
	fmt.Println(3/2) 输出：1
	fmt.Println(3/2.0) 输出：1.5

常量（常量名大写，const关键字，无明确数据类型）
	const (
	PI = 3.1415926
	E  = 2.0
	D  = 9.8
	)
	const (
	A = iota //初台值为0开始
	B        //值为A+1
	C        //值为B+1
	)
	const (
		HELLO = 'HELLO'
	)

GO是强类型语言，类型转换需要显式进行转换，字符串跟数字需要借助函数strconv转换

var s string
	s = strconv.Itoa(n) //int转换成字符串
	fmt.Println(s)
	n, err := strconv.Atoi(s) //字符串转换为int
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

if & case语句

if () {
	
}else if () {
	
}else {
	
}

switch xxx :
case 1: xxx
case 2: xxx
default: xxx

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	n2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	switch os.Args[2] {
	case "+":
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
	}
}

for循环

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
	//for循环3（死循环）用于退出条件不明确
	for {

	}
	//for循环4, 遍历数组元素
	for i, arg := range os.Args {
		fmt.Println(i, args)
	}


文件操作：
写入文件：
	顺序写入
	随机写入
import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("a.txt") //创建文件，如果文件不存在，会创建，文件存在，会清空文件
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello\n") //写入字符串
	f.Close()                //关闭文件
}

//打印字符串并写入文件
f, err := os.Create("fmt.txt")  //创建文件
if err != nil {
	log.Fatal(err)
}
fmt.Fprint(f, "hello")  //打印到控制台并写入文件
fmt.Fprintln(f, "helloln")//打印到控制台并写入文件
s := "hello"
n := 4
fmt.Fprintf(f, "my string is :%s n=%d\n", s, n) //打印到控制台并写入文件
f.Close()

作业：
1. 补全四则运算
2. 反转字符串
3. 实现ps命令