package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gorhill/cronexpr"
)

var d string

func register() {
	flag.StringVar(&d, "d", "2017, 11, 06, 0, 0, 0, 0, time.Local", "Specify date")
	flag.Parse()
}

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
	register()
	stdin := bufio.NewScanner(os.Stdin)
	crontab := getCrontab(stdin)
	nextTime := getNextTime(crontab)

	for _, v := range nextTime {
		fmt.Println(v)
	}
}
