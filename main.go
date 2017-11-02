package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gorhill/cronexpr"
)

func getCrontab(stdin *bufio.Scanner) []string {
	var crontab []string
	for stdin.Scan() {
		crontab = append(crontab, stdin.Text())
	}
	return crontab
}

func getNextTime(crontab []string) []time.Time {
	var nextTime []time.Time
	for _, v := range crontab {
		nextTime = append(nextTime, cronexpr.MustParse(v).Next(time.Now()))
	}
	return nextTime
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	crontab := getCrontab(stdin)
	nextTime := getNextTime(crontab)

	for _, v := range nextTime {
		fmt.Println(v)
	}
}
