//结构体，一种数据的抽象方式
package main

import (
	"fmt"
)

type Student struct { //声明结构体，type + 结构体名称 + struct
	Id   int
	Name string
}

func main() {
	var s Student            //声明一个变量s,类型为student
	var arr [3]Student       //声明一个结构体数组
	var ss []Student         //声明一个结构体切片
	var m map[string]Student //声明一个结构体map
	s.Id = 1                 //初始化变量方式1
	s.Name = "jack"
	fmt.Println(s) //{1 jack}

	s1 := Student{ //初始化变量方式2
		Id:   2,
		Name: "alice",
	}
	fmt.Println(s1) //{2 alice}
	s1 = s
	fmt.Println(s1) //{1 jack}

	s3 := Student{
		Id:   3,
		Name: "Lilei",
	}

	//结构体指针
	var p *Student
	p = &s3
	p.Id = 2
	fmt.Println(s3) //{2 Lilei}

	var p1 *int
	p1 = &s3.Id
	*p1 = 3
	fmt.Println(s3) //{3 Lilei}
}
