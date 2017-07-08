package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")  //调用系统命令
	out, err := cmd.CombinedOutput() //读取标准输出和标准错误，CombinedOutput：标准输出和标准错误混合一起
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
