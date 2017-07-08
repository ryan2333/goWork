package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	f, err := os.Create("ls.out")
	//f, err := os.OpenFile("/dev/null", os.O_WRONLY, 0755) 将标准输出扔到/dev/null
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = f //将标准输出写入到文件
	cmd.Start()
	cmd.Wait()
	//cmd.Start() 和cmd.Wait() 等价于cmd.run()
}
