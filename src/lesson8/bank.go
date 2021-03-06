package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	flag  sync.Mutex //锁，互斥性;互斥锁
	money int
}

func (a *Account) DoPrepare() {
	time.Sleep(time.Millisecond)
}

func (a *Account) GetGongZi(n int) {
	a.money += n
}
func (a *Account) GiveWife(n int) {
	a.flag.Lock()         //使用前先锁定
	defer a.flag.Unlock() //使用完解锁
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Account) Buy(n int) {
	a.flag.Lock()
	defer a.flag.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Account) Left() int {
	return a.money
}

func main() {
	var account Account
	//定义channel
	c := make(chan int, 2) //使用channel同步阻塞
	account.GetGongZi(10)
	go func() {
		account.GiveWife(6)
		c <- 0
	}()
	go func() {
		account.Buy(5)
		c <- 0
	}()
	cnt := 0
	for {
		<-c
		cnt++
		if cnt == 2 {
			break
		}
	}
	fmt.Println(account.Left())
}
