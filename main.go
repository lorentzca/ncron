package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	var crontab []string
	for stdin.Scan() {
		crontab = append(crontab, stdin.Text())
	}

	var nextTime []time.Time
	for _, v := range crontab {
		nextTime = append(nextTime, cronexpr.MustParse(v).Next(time.Now()))
	}

	for _, v := range nextTime {
		fmt.Println(v)
	}
}
