package main

import "fmt"

func main() {
	fmt.Println(dayOfProgrammer(1918))
}

func dayOfProgrammer(year int32) string {
	// Write your code here

	julian := year >= 1700 && year <= 1917
	julian_bisiesto := year%4 == 0

	gregorian := year >= 1919
	gregoria_bisiesto := (year%400 == 0) || (year%4 == 0 && year%100 != 0)

	es_1918 := year == 1918

	var out string

	if julian && julian_bisiesto {
		out = fmt.Sprintf("12.09.%d", year)
	} else if gregorian && gregoria_bisiesto {
		out = fmt.Sprintf("12.09.%d", year)
	} else if es_1918 {
		out = fmt.Sprintf("27.09.%d", year)
	} else {
		out = fmt.Sprintf("13.09.%d", year)
	}

	return out

}
