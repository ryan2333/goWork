package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second) //go 定时器
	cnt := 0
	for _ = range timer.C {
		cnt++
		if cnt > 10 {
			timer.Stop()
			return
		}
		fmt.Printf("%s%d\n", "hello", cnt)
	}
}
