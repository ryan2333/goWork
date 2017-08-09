package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//给定一个url，返回url的status
func printUrl(s string) {
	url := s
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	// st := strings.Fields(resp.Status)[1]
	// fmt.Println(url, st)
	fmt.Println(url, resp.Status)
	// return url, resp.StatusCode, nil
}

func work(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// url := <-ch   //worker协程从channel里面获取url，之后调用printUrl打印url
	// printUrl(url) //调用printUrl打印url

	//方法一：
	// for { //循环读取channel里url
	// 	url, ok := <-ch
	// 	if !ok { //如果ok为false,则跳出循环
	// 		break
	// 	}
	// 	printUrl(url)
	// }

	//方法二：
	for url := range ch {
		printUrl(url)
	}
}

//1. 只要不close可以永远发送数据和接受数据
//2. 如果channel里面没有数据，接收方会阻塞
//3. 如果没有人正在等待channel的数据，发送方会阻塞
//4. 从一个close的channel取数据永远不会阻塞，同时获取的是默认值

//主协程启动一个work协程，同时传递一个channel
//主协程向channel里发送一个url
//work协程从channel里获取url,之后调用printUrl打印url

//启动3个协程
//主协程向channel里发送多个url， 发送完毕后关闭
//worker协程从channel里面获取url，之后调用printUrl打印url
//worker协和不停重复上一条，直到channel关闭

//创建一个waitGroup
//调用Add
//调用Wait等待work协程结束
func main() {
	wg := new(sync.WaitGroup) //创建一个waitGroup
	//	wg.Add(3)                 //Add的次数必须与协程数相同，否则wait不会结束，一般先Add
	ch := make(chan string)
	// go work(ch)  //启动一个协程，传送一个channel
	// url := "http://www.baidu.com"
	// ch <- url  //向channel里传递一个url
	// time.Sleep(3 * time.Second)

	for i := 0; i < 3; i++ { //启动3个协程,并传递chan, waitgroup
		wg.Add(1) //或者采用这种方式，起几个协程，Add几个
		go work(ch, wg)
	}
	urls := []string{"http://www.baidu.com", "http://www.sina.com.cn", "http://www.sohu.com", "http://www.lianjia.com"}
	for _, val := range urls { //主协程向channel里发送多个url， 发送完毕后关闭channel
		ch <- val
	}
	close(ch) //关闭协程
	// time.Sleep(3 * time.Second) //永远不要用sleep方式来进行协程同步
	wg.Wait()

}
