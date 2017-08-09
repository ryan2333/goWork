package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func cleanUrls(u string, urls []string) []string {
	var links []string
	linkState, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
	}
	scheme := linkState.Scheme
	host := linkState.Host
	path := strings.Split(linkState.Path, "/")[1]
	for _, link := range urls {
		if strings.HasPrefix(link, "//") {
			link = scheme + ":" + link
		} else if strings.HasPrefix(link, "/") {
			link = scheme + "://" + host + link
		} else if strings.HasPrefix(link, "http") {

		} else {
			link = scheme + "://" + host + "/" + path + "/" + link
		}
		links = append(links, link)
	}
	return links
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url) //访问url
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { //获取返回code
		return nil, errors.New(resp.Status)
	}
	//	io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp) //goquery生成文档
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) { //查找所有img标签的src属性
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not found")
		}
	})
	return urls, nil
}

func downLoadImgs(urls []string, dir string) error {
	for _, link := range urls {
		resp, err := http.Get(link)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		//	f, err := os.OpenFile(dir+"/"+filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}
		filename := path.Base(link)
		f, err := os.Create(dir + "/" + filename) //创建文件
		// data, err := ioutil.ReadAll(resp.Body)
		// f.Write(data)
		if err != nil {
			return err
		}
		io.Copy(f, resp.Body) //将resp.body内容写入文件
		f.Close()

	}
	return nil
}

func makeTar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		//写入tar header
		//以读取方式打开文件
		//判断目录和文件，如果是文件
		//把文件内容写入到body
		header, err := tar.FileInfoHeader(info, "") //读取头文件信息

		if err != nil {
			return err
		}
		p, _ := filepath.Rel(dir, name)
		// header.Name = name      //替换Name，带全路径的
		header.Name = filepath.Join(basedir, p)
		err = tr.WriteHeader(header) //写入头文件信息
		if err != nil {
			return err
		}
		if info.IsDir() {

		}

		f, err := os.Open(name) //打开文件
		if err != nil {
			return err
		}
		io.Copy(tr, f) //写入文件内容

		f.Close()

		return nil
	})
	return nil
}

func main() {
	fmt.Println(time.Now())
	// link := "http://daily.zhihu.com/"
	// link := os.Args[1]
	link := "http://pic.netbian.com/4kmingxing/index.html"
	urls, err := fetch(link)
	if err != nil {
		log.Fatal(err)
	}
	links := cleanUrls(link, urls)
	// for _, link := range links {
	// 	fmt.Println(link)
	// }
	// tmpdir, err := ioutil.TempDir("", "spider") //创建临时目录
	// if err != nil {
	// 	fmt.Println(err)
	// 	// continue
	// }
	// fmt.Println(tmpdir)
	// defer os.RemoveAll(tmpdir)
	dir := "/Users/yhzhao/Downloads/pics"
	// dir := "G:\\Picture\\pics"
	// dir := os.Args[2]
	err = downLoadImgs(links, dir)
	fmt.Println(time.Now())
	tr, err := os.Create("img.tar.gz")
	// dir := "G:\\Picture\\pics"
	// tr, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer tr.Close()
	// makeTar("/Users/yhzhao/Downloads/pics", tr)
	// makeTar(dir, tr)
	makeTar(dir, tr)
}
