package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.After(3 * time.Second) //几秒后向channel发送信息
	<-c
	fmt.Println("done")
}
