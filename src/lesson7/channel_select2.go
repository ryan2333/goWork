package main

import (
	"fmt"
	"time"
)

func main() {
	// tick := time.NewTicker(1000 * time.Millisecond).C
	tick1 := time.Tick(1000 * time.Millisecond) //time.NewTicker().C的简写方式
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick1:
			fmt.Println("dida...")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("each noodles...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
