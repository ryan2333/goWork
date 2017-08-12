package main

import (
	"crypto/md5"
	"crypto/rc4"
	"io"
	"log"
	"net"
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

func main() {
	remote, err := net.Dial("tcp", "")
	if err != nil {
		log.Fatal(err)
	}

	w := NewCryptoWriter(remote, "123456")
	w.Write([]byte("hello"))

	r := NewCryptoReader(remote, "123456")
	buf := make([]byte, 1024)
	r.Read(buf)
}
