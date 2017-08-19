package main

import (
	"flag"
	"fmt"
	"lesson12/monitor/common"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var (
	transaddr = flag.String("trans", "59.110.12.72:6000", "transfer address")
)

func main() {
	flag.Parse()
	// mem, err := mem.Percent(time.Second)
	sender := NewSender(*transaddr)
	// sender := sender.NewSender(*transaddr)
	go sender.Start()
	ch := sender.Channel()
	ticker := time.NewTicker(time.Second * 5)
	// conn.Write(buf)
	// conn.Write([]byte("\n"))
	// fmt.Println(string(buf))
	// time.Sleep(5 * time.Second)
	for range ticker.C {
		hostname, _ := os.Hostname()
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		metric := &common.Metric{
			Metric:    "cpu.usage",
			Endpoint:  hostname,
			Value:     cpus[0],
			Timestamp: time.Now().Unix(),
		}
		fmt.Println(metric)
		ch <- metric
	}
}
