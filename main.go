package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gorhill/cronexpr"
)

func getCrontab(stdin *bufio.Scanner) []string {
	var c []string
	for stdin.Scan() {
		c = append(c, stdin.Text())
	}
	return c
}

func getNextTime(crontab []string) []time.Time {
	var n []time.Time
	for _, v := range crontab {
		n = append(n, cronexpr.MustParse(v).Next(time.Now()))
	}
	return n
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	crontab := getCrontab(stdin)
	nextTime := getNextTime(crontab)

	for _, v := range nextTime {
		fmt.Println(v)
	}
}
