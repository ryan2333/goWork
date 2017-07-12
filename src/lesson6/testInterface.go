package main

import "fmt"

type Point struct {
	X, Y float64
}

type Path []Point

func (p Path) Instance() float64 {
	return 3.333 //p实现了接口的Instance方法
}

type IInstance interface {
	Instance() float64 //定义接口，包含一个Instance方法
}

func main() {
	var i IInstance
	p := Path{{1, 2}, {3, 4}}
	i = p //只有p里包含Instance方法后，才能相等
	fmt.Println(i)
}
