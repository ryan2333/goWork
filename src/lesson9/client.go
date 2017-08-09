package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:7777"
	conn, err := net.Dial("tcp", addr) //建立tcp连接,返回一个连接和错误信息
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String())                //打印连接的远端地址和端口
	fmt.Println(conn.LocalAddr().String())                 //打印连接的本地地址和端口
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")) //向远端发送数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("write size: ", n) //向远端发送了多少字节

	//获取远端内容，按字节读取
	// buf := make([]byte, 10)
	// n, err = conn.Read(buf) //读取远端数据,n表示读取多少字节

	// if err != nil && err != io.EOF {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(buf[:n]))  //打印从远端读取的内容
	// fmt.Println("Read size: ", n) //从远端读取了多少字节
	// for {                         //每次读取10字节，循环读取
	// 	n, err := conn.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(string(buf[:n]))
	// }

	// //获取远端内容，按行读取
	// content := bufio.NewReader(conn)
	// for {
	// 	line, err := content.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(line)
	// }

	//使用io.copy读取
	io.Copy(os.Stdout, conn)

	conn.Close()
}
