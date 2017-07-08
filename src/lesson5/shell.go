package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()                     //获取主机名
	prompt := fmt.Sprintf("[zhaoyh@%s]$ ", host) //显示提示符
	r := bufio.NewScanner(os.Stdin)              //读取标准输入，scanner按行读取
	for {
		fmt.Print(prompt)
		if !r.Scan() { //读取一行，返回True or False,True代表读取一整行，False表示没有数据了
			break
		}
		line := r.Text() //将scan读取的数据拿出来，赋值给line
		/*
			等价于：
			line, _:= r.ReadString('\n')
			line = strings.TrimSpace(line)

		*/
		if len(line) == 0 {
			continue
		}
		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin //连接系统的标准输入、输出、和错误到程序的标准输入、输出和错误
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
