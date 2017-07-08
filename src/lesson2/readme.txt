课程内容：

同一个目录下的go文件只能是一个package, 同一个目录下go文件中变量，函数是共享的
main package

在一个目录下，go build不加任何参数，将会将当前目录所有build为一个包
	
标准go程序结构：
package main  //引入包，放在程序的第一行，两种package，一种是库package, 一种是二进制package;二进制package使用main表示，库package的名字跟go文件所在的目录名一样
	//整个程序只能存在一个package main
	//三方库文档： godoc.org/xxx

import (   //引入第三方库,引入多个库用圆括号
	"fmt"   //引入三方库要加引号，必须引入全路径，路径是以$GOPATH/src为根路径
	"os"   //跟系统交互的包
)

func main(){   //声明一个函数，main函数是二进制程序的入口，main函数结束了整个程序就结束了
	var s, sep string  //初台化变量s、sep，变量默认为空值
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]  //将参数拼接为一个字符串
		sep = " "   //sep赋值为” “，以空格拼接字符串 
	}
	fmt.Println(s)
}


go run 跟go build 或go install区别
go run针对单个文件
go build或go install 是针对package

引入github库
1. go get github.com/xxx/xxx   //download 库
2. import "github.com/xxx/xxx"   //导入库
3. godoc.org/github.com/xxx/xxx //查看文档

代码风格：
所有的代码只有一种,gofmt的风格
gofmt -w xx.go


变量声明：
var x int
var y string = "hello"
var x, y int
var (   //声明全局变量是使用
	x int
	y int
	z string
)

//变量初始化零值，安全
package main

import "fmt"

func main() {
	var (
		x int
		y float32
		z string
		p *int
		l bool
	)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", p)
}

短变量   //一般用于局部变量声明
i := 0   //int, 0默认为int类型
s := "hello"   //string类型
i, j := 0, 1  //批量初始化，类似于python元组

指针
*T 即为类型T的指针， T可以为任何类型，包括指针类型，**T，表示指针类型的指针
&t 即为取变量t的地址
*p 即为取指针变量所指向的内容
func main() {
	var x int = 5
	var p *int  //声明指针，类型为int， 指针的类型必须跟经引用的变量类型一致
	p = &x  //只能将合法变量的内存地址赋给指针
	fmt.Println(p)   //打印指针内存地址
	fmt.Println(*p)   //打印指针内存地址指向变量的值
}

flag包的使用
package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")  //第一个是命令行参数，即-s;第二个是默认值，第三个是描述

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))  
}


变量的生命周期
#垃圾回收  
#栈和堆 栈：函数运行结束后内存释放，堆：程序结束后内存释放
#逃逸分析机制


多变量赋值：
x,y = y,x
func fib(n int) int {
	for i := 0; i < n; i++ {
		x,y = y, x + y
	}
	return x
}


