package main

import (
	"fmt"
	"strings"
)

func main() {

	var steps int32
	var path string

	steps = 8
	path = "UDDDUDUU"
	fmt.Println(countingValleys(steps, path))
}

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	//O mejor dicho: contar pozos

	m := strings.Split(path, "")

	var countPozo int32
	var pozo int32

	for _, p := range m {
		if string(p) == "U" {
			pozo++
			if pozo == 0 {
				countPozo++
			}
		} else if string(p) == "D" {
			pozo--
		}
	}
	return countPozo
}
