package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	//s := os.Args[1]
	s := "http://59.110.12.72:7070/golang-spider/img.html"
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("scheme", u.Scheme) //http://xxx.xxx.com/path/a.jpg?a=b&c=d#ssss
	fmt.Println("host", u.Host)
	fmt.Println("path", u.Path)
	fmt.Println(u.EscapedPath())
	fmt.Println("queryString", u.RawQuery) //查询字串
	fmt.Println("user", u.User)            //ftp://user:password@xxx.com
	fmt.Println("xxx", u.Fragment)         //锚点
}
