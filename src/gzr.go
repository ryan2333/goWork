package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t := time.Now().UTC().String()
	ts := strings.Fields(t)[0]
	ta := strings.Split(ts, "-")
	for _, i := range ta {
		fmt.Printf("%T, %v\n", i, i)
	}

}
