//变量声明练习
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
	i := 0
	s := "hello"
	j, k := 0, 1
	fmt.Printf("x=%v\n", x)
	fmt.Printf("y=%v\n", y)
	fmt.Printf("z=%v\n", z)
	fmt.Printf("l=%v\n", l)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("i=%v\n", i)
	fmt.Printf("s=%v\n", s)
	fmt.Printf("j=%v\n", j)
	fmt.Printf("k=%v\n", k)
}
