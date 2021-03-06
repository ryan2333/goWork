package main

import (
	"bufio"
	"crypto/md5"
	"crypto/rc4"
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}
}

//把b里面的数据进行加密，之后写入w.w里面
//调用w.w.Write进行写入
func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf, b)
	return w.w.Write(buf)
}

func NewCryptoReader(r io.Reader, key string) io.Reader {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)
	}
	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}
}

func (r *CryptoReader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err
}

var (
	target = flag.String("target", "127.0.0.1:7778", "target host")
)

func handshake(r *bufio.Reader, conn net.Conn) error {
	//获取客户端协议版本，1个字节，ReadByte表示读取1个字节
	version, _ := r.ReadByte()
	// version := mustReadByte(r)
	log.Printf("version:%d\n", version)
	if version != 5 {
		return errors.New("bad version")
	}
	//读取认证方式长度，1个字节
	nmethods, _ := r.ReadByte()
	// nmethods := mustReadByte(r)
	log.Printf("nmethods:%d\n", nmethods)

	//读取认证方式，根据认证方式长度，创建对应的BUF，当BUF读满则完整读出认证方式
	buf := make([]byte, nmethods)
	io.ReadFull(r, buf) //从r中读取内容，直到buf填充满
	log.Printf("buf: %v", buf)

	//给客户端返回服务端支持的认证方式及长度
	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

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
		r := NewCryptoReader(conn, "123456")
		io.Copy(toDst, r)
		toDst.Close()
	}()

	//go 从目标服务器发送到客户端的协程
	go func() {
		defer wg.Done()
		w := NewCryptoWriter(conn, "123456")
		io.Copy(w, toDst)
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
