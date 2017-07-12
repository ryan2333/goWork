package main

import (
	"fmt"
	"math"
)

type Point struct { //type用来定义新类型
	X, Y float64
}

func (p Point) Distance(q Point) float64 { //括号内要为哪个结构体绑定方法
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point //自定义类型Path, 类型为Point结构体组成的切片

func Distance(path Path) float64 { //将参数path，类型Path传入函数，求切片path中所有点之间的路径长度总和
	var sum1 float64
	for i := 0; i < len(path)-1; i++ {
		p := path[i]
		q := path[i+1]
		sum1 += p.Distance(q)
	}
	return sum1
}

func (path Path) Distance() float64 { //将Distance方法绑定到path上，求切片path中所有点之间的路径长度总和
	var sum float64
	for i := 0; i < len(path)-1; i++ {
		p := path[i]
		q := path[i+1]
		sum += p.Distance(q)
	}
	return sum
}

func main() {
	path1 := []Point{{1, 2}, {3, 4}, {5, 6}}
	path := Path{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(Distance(path1)) //函数求路径长度
	fmt.Println(path.Distance()) //方法求路径长度
}
