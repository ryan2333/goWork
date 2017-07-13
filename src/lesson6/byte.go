//统计字符
package main

import (
	"fmt"
	"io"
	"os"
)

// type ByteCounter struct {
// 	Sum int
// }

type ByteCounter int //方式一

type ByteCounter1 struct { //方式二
	Sum int
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p)) //方式一
	return len(p), nil
}

func (b *ByteCounter1) Write(p []byte) (int, error) {
	b.Sum += len(p) //方式二
	return len(p), nil
}

func main() {
	// b := new(ByteCounter) //方式一
	// io.Copy(b, os.Stdin)
	// fmt.Println(*b)
	c := new(ByteCounter1) //方式二
	io.Copy(c, os.Stdin)
	fmt.Println(c.Sum)
}
