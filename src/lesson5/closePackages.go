package main

import (
	"errors"
	"fmt"
)

func addn(n int) func(int) int {
	return func(m int) int { //此匿名函数与参数n，组成一个闭包
		return m + n
	}
}

func iter(s []int) func() (int, error) { //python生成器
	var i = 0
	return func() (int, error) { //多返回值
		if i >= len(s) {
			return 0, errors.New("end")
		}
		n := s[i]
		i += 1
		return n, nil
	}
}

func main() {
	// f := addn(3)，此时参数n已确定值，

	// fmt.Println(f(18)) 当调用f函数时，m值才确定
	f := iter([]int{1, 2, 3})
	for {
		n, err := f()
		if err != nil {
			break
		}
		fmt.Println(n)
	}
}
