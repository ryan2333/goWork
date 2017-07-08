//文件读取的几种方式
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	v2()
}

func v2() {
	var f *os.File
	var err error
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1]) //如果输入了文件名参数，则打开文件
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		f = os.Stdin //如果没有输入文件名参数，则从标准输入里读取
	}
	io.Copy(os.Stdin, f) //v2版本，将标准输入f对象
}

func v1() {
	var f *os.File
	var err error
	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1]) //如果输入了文件名参数，则打开文件
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		f = os.Stdin //如果没有输入文件名参数，则从标准输入里读取
	}
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			return
		}
		os.Stdout.Write(buf[:n])
	}
}

func openFileMothod() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := make([]byte, 4096) //Read，按块读取，效率较低，裸读取，很少使用
	n, err := f.Read(buf)
	fmt.Println(buf[:n])

	r := bufio.NewReader(f) //bufio 加上buffer的读取，高效的读取方式
	r.ReadByte()            //按字节读取
	r.Read(buf)             //按字节读取，每次读取多少字节
	r.ReadString('\n')      //按行读取

	r1 := bufio.NewScanner(f) //按行读取，按分隔符读取

	ioutil.ReadFile("a.txt") //小文件一次性读取，不打开文件
	ioutil.ReadAll(f)        //一次性将文件传进来，文件打开后

	io.Copy(dFile, sFile) //神器， io.copy操作文件的高级方法
}

func pipe() {
	line := "ls |grep file"
	cmds := strings.Split(line, "|") //经竖线分隔两个命令
	s1 := strings.Fields(cmds[0])    //将第一个命令以空格分成一个切片
	s2 := strings.Fields(cmds[1])    //将第二个命令以空格分成一个切片

	r, w := io.Pipe()
	cmd1 := exec.Command(s1[0], s1[1:]...) //第一个命令
	cmd2 := exec.Command(s2[0], s2[1:]...) //第二个命令
	cmd1.Stdin = os.Stdin                  //读取标准输入
	cmd1.Stdout = w                        //输出到管道
	cmd2.Stdin = r                         //读到cmd1的输出为命2的输入
	cmd2.Stdout = os.Stdout                //输出到标准输出
	cmd1.Start()                           //运行cmd1命令
	cmd2.Start()
	//cmd2.Wait()
	cmd1.Wait()
}
