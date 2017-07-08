package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("a.txt") //创建文件
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("hello\n") //写入字符串
	f.Close()                //关闭文件

	fh, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644) //第一个参数：文件名，第二个参数，存在即追加，不存在即李建，第三个参数为权限
	//os.O_APPEND 追加 os.O_CREAT 不存在即创建 os.O_TRUNC每次打开都清空内容
	if err != nil {
		log.Fatal(err)
	}
	fh.WriteString("hello\n") //写入字符串
	fh.Close()

	ff, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644) //第一个参数：文件名，第二个参数，存在即追加，不存在即李建，第三个参数为权限
	//os.O_APPEND 追加 os.O_CREAT 不存在即创建
	if err != nil {
		log.Fatal(err)
	}
	ff.Seek(3, os.SEEK_SET) //第起始位置，偏移3个字符开始写入SEEK_SET表示起始位置，SEEK_CUR表示当前位置,SEEK_END表示结束位置
	ff.WriteString("SSSSS")
	ff.Seek(3, os.SEEK_END)
	ff.WriteString("EEEE")
	ff.Close()
}
