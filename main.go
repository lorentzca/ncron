package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

func cStdin(stdin *bufio.Scanner) []string {
	var s []string

	for stdin.Scan() {
		s = append(s, stdin.Text())
	}
	return s
}

func crontabMap(s []string) map[string]string {
	m := map[string]string{}

	for _, v := range s {
		t := strings.Join(strings.Split(v, " ")[0:5], " ")
		j := strings.Join(strings.Split(v, " ")[5:], " ")
		m[t] = j
	}
	return m
}

func getNextTimeSpecifiedDate(crontabMap map[string]string) []time.Time {
	var n []time.Time

	for k, _ := range crontabMap {
		n = append(n, cronexpr.MustParse(k).Next(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)))
	}
	return n
}

func main() {
	register()
	stdin := bufio.NewScanner(os.Stdin)
	stdinSlice := cStdin(stdin)
	cMap := crontabMap(stdinSlice)

	nextTimeSpecified := getNextTimeSpecifiedDate(cMap)

	for _, v := range nextTimeSpecified {
		fmt.Println(v)
	}
}
