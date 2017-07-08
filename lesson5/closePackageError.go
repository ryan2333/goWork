//闭包坑
package main

import (
	"fmt"
)

func main() {
	var flist []func()
	for i := 0; i < 3; i++ {
		j := i                         //加上这行，破解闭包坑
		flist = append(flist, func() { //或者在匿名函数中传入参数i
			fmt.Println(j)
		})
	}
	for _, f := range flist {
		f()
	}
}
