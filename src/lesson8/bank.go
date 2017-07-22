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
	account.GetGongZi(10)
	go account.GiveWife(6)
	go account.Buy(5)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(account.Left())
}
