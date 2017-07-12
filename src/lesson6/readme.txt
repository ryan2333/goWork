总览
	对象和方法
	接口使用
	go里面的unix哲学


type是用来定义类型的
不同的对象可以有相同的方法名

package main

type Point struct{  //定义Point类型为结构体
	X,Y float64
}

type Path []Point  //定义Path类型为结构体组成的切片

func (path Path) Distance() float64{
	//求结构体切片的路径长度
}

想为一个类型绑定方法，需要使用type声明一个新的类型


可见性
	通过首字母大小写来控制可见性
	可见性是package级别

声明接口
	
	type Writer interface {
		Write(p []byte)(n int, err error)
	}