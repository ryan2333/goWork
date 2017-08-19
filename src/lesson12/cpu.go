package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

func main() {
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(cpus[0], cpus)
	loadavg, err := load.Avg()
	if err != nil {
		panic(err)
	}
	fmt.Println(loadavg)
}
