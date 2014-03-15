package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	datafile := filepath.Join(os.Getenv("HOME"), ".cchecker", "ccheck.txt")
	f, err := os.Open(datafile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()
	now := time.Now()
	today := fmt.Sprintf("%02d.%02d", now.Day(), now.Month())
	reader := bufio.NewReader(f)

	var line string
	for err != io.EOF {
		line, err = reader.ReadString('\n')
		if strings.HasPrefix(line, today) {
			fmt.Println(strings.Replace(line, "\\n", "\n      ", -1))
		} else if strings.HasPrefix(line, "[") {
			m, d, i := ParseRelDate(line)
			c := WeekdayIndex(m, d, i)
			if c.Month() == now.Month() && c.Day() == now.Day() || i == now.YearDay() {
				fmt.Println(strings.Replace(line, "\\n", "\n"+strings.Repeat(" ", strings.Index(line, "]")+2), -1))
			}
		}
	}
}

// WeekdayIndex returns Date of n'th Weekday of given Month
func WeekdayIndex(m time.Month, wd time.Weekday, n int) time.Time {
	date := time.Date(time.Now().Year(), m, time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	FirstDayOfMonth := date.AddDate(0, 0, -date.Day()+1)

	var cnt int
	for d := FirstDayOfMonth; d.Month() == date.Month(); d = d.Add(24 * time.Hour) {
		if d.Weekday() == wd {
			cnt = cnt + 1
		}
		if cnt == n {
			return d
		}
	}
	return time.Time{}

}
// ParseRelDate parses relative dates like "3rd Sunday of March"
func ParseRelDate(str string) (m time.Month, d time.Weekday, index int) {
	str = strings.ToLower(str)
	re, _ := regexp.Compile(`(\d+)`)
	data := re.FindString(str)
	index, _ = strconv.Atoi(data)
	switch {
	case strings.Contains(str, "first"):
		index = 1
	case strings.Contains(str, "second"):
		index = 2
	case strings.Contains(str, "third"):
		index = 3
	case strings.Contains(str, "forth"):
		index = 4
	case strings.Contains(str, "fifth"):
		index = 5
	case strings.Contains(str, "sixth"):
		index = 6
	}
	switch { // <weekday>
	case strings.Contains(str, "sun"):
		d = time.Sunday
	case strings.Contains(str, "mon"):
		d = time.Monday
	case strings.Contains(str, "tue"):
		d = time.Tuesday
	case strings.Contains(str, "wed"):
		d = time.Wednesday
	case strings.Contains(str, "thu"):
		d = time.Thursday
	case strings.Contains(str, "fri"):
		d = time.Friday
	case strings.Contains(str, "sat"):
		d = time.Saturday
	default:
		d = time.Now().Weekday()
	}
	switch {
	case strings.Contains(str, "jan"):
		m = time.January
	case strings.Contains(str, "feb"):
		m = time.February
	case strings.Contains(str, "mar"):
		m = time.March
	case strings.Contains(str, "apr"):
		m = time.April
	case strings.Contains(str, "may"):
		m = time.May
	case strings.Contains(str, "jun"):
		m = time.June
	case strings.Contains(str, "jul"):
		m = time.July
	case strings.Contains(str, "aug"):
		m = time.August
	case strings.Contains(str, "sep"):
		m = time.September
	case strings.Contains(str, "oct"):
		m = time.October
	case strings.Contains(str, "nov"):
		m = time.November
	case strings.Contains(str, "dec"):
		m = time.December
	default:
		m = time.Now().Month()
	}
	return
}

// vim: ts=8:sw=8:et
