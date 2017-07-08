package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//创建map
	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	//or
	ages1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(ages, ages1)

	//获取元素
	aa := ages["a"] + 2
	fmt.Println(ages["a"], aa)
	c, ok := ages["c"] //如果c存在，则OK为true,不存在，ok为false
	if ok {
		fmt.Println(c)
	} else {
		fmt.Println("not fount")
	}

	//删除元素
	delete(ages, "a") //第一个参数是map变量，第二个参数是key
	fmt.Println(ages)

	//map遍历
	for name, age := range ages1 {
		fmt.Println("name", name, "age", age)
	}

	//只遍历名字
	for name := range ages1 {
		fmt.Println("name", name)
	}

	count := make(map[string]int)                                                               //初始化一个map变量count,key为字符串，值为int
	buf, err := ioutil.ReadFile("/Users/yhzhao/Documents/workspace/Go/src/lesson4/mymap/a.txt") //读取文件内容
	if err != nil {
		fmt.Println(err)
		return
	}
	words := strings.Fields(string(buf)) //对读取的字符串，以空格为分割成一个数组
	for _, word := range words {         //循环单词数组
		count[word] += 1 //在count中查找key为单词的值，如果没找到值为0，对值进行加1
	}
	for k, v := range count { //循环count的key和值
		fmt.Printf("单词：%-20s出现次数：%d\n", k, v)
	}

}
