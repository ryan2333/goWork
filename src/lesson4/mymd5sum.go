package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// data := []byte("hello")
	// md5sum := md5.Sum(data)
	// fmt.Printf("%x\n", md5sum)
	if len(os.Args) == 1 {
		fmt.Println("请输入文件名")
		return
	}
	for _, filename := range os.Args[1:] {
		buf, err := ioutil.ReadFile(filename) //读取文件内容
		if err != nil {
			fmt.Println(err)
			return
		}
		data := []byte(buf) //将文件内容以字节方式赋值到数组data
		//		fmt.Println(data)
		md5sum := md5.Sum(data)                  //对数组内容进行md5值计算
		fmt.Printf("%x\t%s\n", md5sum, filename) //打印md5值和文件名
	}
}
