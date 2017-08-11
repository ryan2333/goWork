package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)

func handleConn(conn net.Conn) {
	//建立到目标服务器的连接
	toDst, err := net.Dial("tcp", *target)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//go 从客户端到目标服务器的协程
	go func() {
		defer wg.Done()
		io.Copy(toDst, conn)
		toDst.Close()
	}()

	//go 从目标服务器发送到客户端的协程
	go func() {
		defer wg.Done()
		io.Copy(conn, toDst)
		conn.Close()
	}()
	//等待两个协程结束
	wg.Wait()
}

func main() {
	flag.Parse()
	//建立listen
	addr := ":7777"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		//接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}
