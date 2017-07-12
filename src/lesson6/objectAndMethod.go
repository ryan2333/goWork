package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

//声明函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y) //Hpoy == sqrt((X2-X1)^2 + (Y2-Y1)^2),求两点之间的距离
}

//声明方法,功能等价于上面的函数
//将函数变为方法，将其中一个参数向外提到函数名前面
func (p Point) Distance(q Point) float64 { //括号内要为哪个结构体绑定方法，即给Point绑定Distance方法
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) //"5", function
	fmt.Println(p.Distance(q))  //"5", method
}
