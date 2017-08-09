package main

import (
	"log"
	"net"
)

func handConn(conn net.Conn) { //conn类型为net.conn
	conn.Write([]byte("hello golang\n")) //发送数据
	// time.Sleep(time.Minute)
	conn.Close() //关闭连接
}

func main() {
	addr := ":7777"                          //addr为IP:端口，一般忽略IP
	listener, err := net.Listen("tcp", addr) //监听端口
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //程序退出，关闭监听
	for {                  //循环接收请求
		conn, err := listener.Accept() //接收新的连接，阻塞连接
		if err != nil {
			log.Fatal(err)
		}
		go handConn(conn) //启动处理连接协程，哪里阻塞go哪里

	}
}
