package main

import (
	"fmt"
	"math"
)

type Point struct { //type用来定义新类型
	X, Y float64
}

func (p Point) Distance2Point(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance() float64 {
	return math.Hypot(p.X, p.Y) //sqrt(x^2+y^2)
}

type Path []Point

func (p Path) Distance() float64 {
	var sum float64
	for i := 0; i < len(p)-1; i++ {
		sum += p[i].Distance2Point(p[i+1])
	}
	return sum
}

type IDistance interface {
	Distance() float64
}

func print(p IDistance) { //这里参数设置成接口，可以打印任何数据
	fmt.Println(p.Distance())
}

func main() {
	var path Path           //声明path变量，类型为Path
	path = make([]Point, 3) //声明一个数组，元素类型为point结构体
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 3, Y: 4}
	p3 := Point{X: 5, Y: 6}
	path[0] = p1 //第一个元素
	path[1] = p2 //第二个元素
	path[2] = p3 //第三个元素

	fmt.Println(p1.Distance2Point(p2))
	var i IDistance //声明一个接口变量
	i = p1          //p1有Distance方法，所以可以相等
	fmt.Println(i.Distance())
	i = p2 //p2有Distance方法，所以可以相等
	fmt.Println(i.Distance())
	i = p3 //p3有Distance方法，所以可以相等
	fmt.Println(i.Distance())
	print(path)
	print(p1)
}
