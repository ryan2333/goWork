package main

//求当月地铁票花费总数
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func calcYear(year, month string) (int, error) {
	//计算输入月份或当月的工作日天数
	var y int
	var err error
	var days, workDays int
	if y, err = strconv.Atoi(year); err != nil {
		return 0, err
	}

	if month == "2" {
		if y%4 == 0 && y%100 != 0 || y%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	} else if month == "04" || month == "06" || month == "09" || month == "11" {
		days = 30
	} else {
		days = 31
	}
	start := 1
	var startDay string
	for start <= days {
		if start < 10 {
			startDay = fmt.Sprintf("0%d", start)
		} else {
			startDay = fmt.Sprintf("%d", start)
		}
		cD, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%s-%s", y, month, startDay))
		// t, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%s-%s", y, month, startDay))
		if err != nil {
			log.Println(err)
		}

		if cD.Weekday().String() != "Saturday" && cD.Weekday().String() != "Sunday" {
			// fmt.Printf("days: %s, weekday: %s\n", startDay, cD.Weekday().String())
			workDays += 1
		}
		start += 1
	}
	return workDays, nil
}

func calcSubway(m int) float64 {
	//计算当月地铁总票价
	var sum float64 = 0
	var ticketa float64 = 6
	var ticketm float64 = 6
	startDay := 1
	for startDay <= m {
		if sum < 100 {
			sum += ticketm
			if sum < 100 {
				sum += ticketa
			} else {
				sum += ticketa * 0.8
				fmt.Printf("days: %d, ticket: %s, sum: %.2f\n", startDay, "ticketa", sum)
			}
		} else if sum < 150 {
			sum += ticketm * 0.8
			if sum < 150 {
				sum += ticketa * 0.8
			} else {
				sum += ticketa * 0.5
				fmt.Printf("days: %d, ticket: %s, sum: %.2f\n", startDay, "ticketa", sum)
			}
		} else {
			sum += (ticketa + ticketm) * 0.5
			fmt.Printf("days: %d, sum: %.2f\n", startDay, sum)
		}
		// fmt.Printf("day: %d, ticket: %.2f\n", startDay, sum)
		startDay += 1
	}
	return sum
}

func main() {
	// date := "2017-12-03"
	var date []string
	if len(os.Args) < 2 {
		dt := strings.Fields(time.Now().String())[0]
		date = strings.Split(dt, "-")

	} else {
		dt := os.Args[1]
		date = strings.Split(dt, "-")
	}
	for _, d := range date {
		if _, err := strconv.Atoi(d); err != nil {
			log.Fatal("输入的日期不正确，exp: 2010-11-22")
			os.Exit(0)
		}
	}
	y := date[0]
	m := date[1]
	// d := date[6:8]
	// fmt.Printf("year: %s, Month: %s, date: %s\n", y, m, d)
	days, err := calcYear(y, m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("month: %s, workdays: %d\n", m, days)
	sum := calcSubway(days)
	fmt.Println(sum)
	fmt.Println("five second to exit")
	time.Sleep(5 * time.Second)
}
