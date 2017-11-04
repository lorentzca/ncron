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
	y   int
	mon int
	d   int
	h   int
	min int
	j   bool
)

func register() {
	flag.IntVar(&y, "y", time.Now().Year(), "Specify Year")
	flag.IntVar(&mon, "mon", int(time.Now().Month()), "Specify Month")
	flag.IntVar(&d, "d", time.Now().Day(), "Specify Day")
	flag.IntVar(&h, "h", time.Now().Hour(), "Specify Hour")
	flag.IntVar(&min, "min", time.Now().Minute(), "Specify Minute")
	flag.BoolVar(&j, "j", false, "Show Job")
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
		cj := strings.Join(strings.Split(v, " ")[5:], " ")
		m[t] = cj
	}
	return m
}

func getNextTimeSpecifiedDate(crontabMap map[string]string) map[time.Time]string {
	n := map[time.Time]string{}

	for k, v := range crontabMap {
		t := cronexpr.MustParse(k).Next(time.Date(y, time.Month(mon), d, h, min, 0, 0, time.Local))
		n[t] = v
	}
	return n
}

func main() {
	register()
	stdin := bufio.NewScanner(os.Stdin)
	stdinSlice := cStdin(stdin)
	cMap := crontabMap(stdinSlice)

	nextTimeSpecified := getNextTimeSpecifiedDate(cMap)

	if j {
		for k, v := range nextTimeSpecified {
			fmt.Println(k, v)
		}
	} else {
		for k, _ := range nextTimeSpecified {
			fmt.Println(k)
		}
	}
}
