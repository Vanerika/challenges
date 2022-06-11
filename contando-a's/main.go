package main

import (
	"fmt"
)

func main() {
	s := "abca"
	var n int64
	n = 30

	fmt.Println(repeatedString(s, n))

}

func repeatedString(s string, n int64) int64 {
	// Write your code here
	var count int64

	var entero int64
	entero = n / int64(len(s))
	var resto int64
	resto = n % int64(len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			count++
		}
	}

	count = entero * count

	if resto > 0 {
		s1 := s[:resto]
		for i := 0; i < len(s1); i++ {
			if s[i] == 'a' {
				count++
			}
		}
	}

	return count
}
