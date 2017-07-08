package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := Student{
		Id:   1,
		Name: "alice",
	}
	buf, err := json.Marshal(s) //序列化, 对内存中的对象转换成字符串
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("/Users/yhzhao/Documents/workspace/Go/src/lesson4/Students.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(buf)) //序列化, 将序列化后的字串符写入文件
	f.Close()

	var s1 Student
	f1, err := os.Open("/Users/yhzhao/Documents/workspace/Go/src/lesson4/Students.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f1) //反序列化, 将磁盘中的数据读入内存
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		ferr := json.Unmarshal([]byte(line), &s1) //反序列化，将读出的字符串转换成原始数据
		if ferr != nil {
			log.Fatal(err)
		}
		fmt.Println(s1)
	}
	f1.Close()

}
