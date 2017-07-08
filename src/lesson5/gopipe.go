//交互式调用
package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	out, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	f := bufio.NewReader(out)

	for {
		line, err := f.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
	err := cmd.Wait() //err可以获取命令的退出信息，通过wait方法可以获取子进程的退出信息
}
