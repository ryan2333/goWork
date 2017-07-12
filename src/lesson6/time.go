package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration //定义一个变量n,类型是Duration,表示一个固定的时间长度
	n = time.Hour       //此处可以为Hour(小时), second(秒), minute(分), microsecond(微秒), millisecond(毫秒)
	fmt.Println(time.Now())
	//	time.Sleep(3 * n) //表示sleep3杪
	fmt.Println("Time now: ", time.Now())
	fmt.Println("n int64: ", int64(n))
	fmt.Println("n String: ", n.String())
	n = 3*time.Hour + 30*time.Minute
	fmt.Println("Hours: ", n.Hours())     //按小时显示
	fmt.Println("Minutes: ", n.Seconds()) //按秒显示
	fmt.Println("Secount: ", n.Minutes()) //按分钟显示

	t := time.Now()                      //表示现在时间
	t1 := t.Add(-time.Hour)              //表示1小时前
	fmt.Println("one hour ago: ", t1)    //打印1小时前
	fmt.Println("time sub: ", t.Sub(t1)) //计算两个时间之间的差
}
