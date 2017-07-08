package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func read(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024) //读取文件，每次读取1024字节
	for {
		n, err := f.Read(buf)
		if err == io.EOF { //EOF error
			break
		}
		if err != nil {
			return "", err
		}
		total = append(total, buf[:n]...)
	}
	return string(total), nil
}

func readR(f *os.File) (string, error) {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err //错误处理方式一，直接退出
	}
	return string(buf), nil
}

func main() {
	f, err := os.Open("readme.txt")
	if err != nil {
		log.Fatalf("open error: %v", err) //错误处理方式二，打印错误
	}

	var content string //错误处理方式三，重试
	retries := 3
	for i := 1; i <= retries; i++ {
		content, err = readR{f}
		if err == nil {
			break
		}
		time.Sleep(time.Second << i)
	}
	fmt.Println(content)

	s, err := read(f)
	if err != nil {
		log.Fatalf("read error: %v", err) //错误处理方式二，打印错误
	}
	fmt.Println(s)
}
