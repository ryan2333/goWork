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
	//定义waitgroup
	wg := new(sync.WaitGroup)
	wg.Add(2)
	account.GetGongZi(10)
	go func() {
		account.GiveWife(6)
		wg.Done()
	}()
	go func() {
		account.Buy(5)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(account.Left())
}
