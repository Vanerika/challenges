package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(timeConversion("07:05:45PM"))

}

func timeConversion(s string) string {
	// Write your code here

	s = strings.ToUpper(s)

	s1 := strings.Split(s[:8], ":")

	hour, _ := strconv.Atoi(s1[0])
	is_pm := strings.HasSuffix(s, "PM")
	is_am := strings.HasSuffix(s, "AM")

	if is_pm && hour != 12 {
		hour = hour + 12
	} else if is_am && hour == 12 {
		hour = 00
	}

	result := fmt.Sprintf("%02v:%v:%v", hour, s1[1], s[2])

	return result

}
