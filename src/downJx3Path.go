package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

func downLoadImgs(ch chan string, wg *sync.WaitGroup, savePath string) error {
	for link := range ch {
		resp, err := http.Get(link)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}
		filename := path.Base(link)
		f, err := os.Create(savePath + "/" + filename) //创建文件
		if err != nil {
			return err
		}
		io.Copy(f, resp.Body) //将resp.body内容写入文件
		f.Close()

	}
	wg.Done()
	return nil
}

func downloadPath(urls []string, savePath string) {
	wg := new(sync.WaitGroup)
	wg.Add(5)
	downUrlCh := make(chan string)
	for i := 0; i < 5; i++ {
		go downLoadImgs(downUrlCh, wg, savePath)
	}
	for _, link := range urls {
		downUrlCh <- link
	}
	close(downUrlCh)

}
func main() {
	if len(os.Args) < 3 {
		fmt.Println("输入的参数不正确，example: downJx3Path 补丁包保存目录 补丁包名称")
		os.Exit(0)
	}
	savePath := os.Args[1]
	url := "http://jx3.autoupdate.kingsoft.com/jx3hd/zhcn_exp/jx3_c_%s_zhcn_exp_patch.exe"
	pathNames := os.Args[2:]
	var urls []string
	for _, v := range pathNames {
		urls = append(urls, fmt.Sprintf(url, v))
	}
	downloadPath(urls, savePath)
}
