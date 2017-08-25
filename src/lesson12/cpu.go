package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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

	memstat, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	fmt.Println(memstat.UsedPercent)

	diskstat, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}
	fmt.Println(diskstat.UsedPercent)
}
