package main

import (
	"fmt"
)

func main() {

	scores := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	fmt.Println(breakingRecords(scores))

}
func breakingRecords(scores []int32) []int32 {
	// Write your code here

	var count_min int32
	var count_max int32

	max := scores[0]
	min := scores[0]

	for _, j := range scores {
		if max < j {
			count_max++
			max = j
		}
		if min > j {
			count_min++
			min = j
		}
	}

	records := []int32{count_max, count_min}
	return records

}
