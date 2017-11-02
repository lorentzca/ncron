package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gorhill/cronexpr"
)

var (
	y int
	m int
	d int
)

func register() {
	flag.IntVar(&y, "y", 1989, "Specify Year")
	flag.IntVar(&m, "m", 7, "Specify Month")
	flag.IntVar(&d, "d", 3, "Specify Day")
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

func getNextTimeSpecifiedDate(crontab []string) []time.Time {
	var n []time.Time

	for _, v := range crontab {
		n = append(n, cronexpr.MustParse(v).Next(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)))
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

	nextTimeSpecified := getNextTimeSpecifiedDate(crontab)
	for _, v := range nextTimeSpecified {
		fmt.Println(v)
	}
}
