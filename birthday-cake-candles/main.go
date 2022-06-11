package main

import (
	"fmt"
	"sort"
)

func main() {

	candles := []int32{1, 3, 5, 7, 9, 9, 9, 9, 9}
	fmt.Println(birthdayCakeCandles(candles))

}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here

	sort.Slice(candles, func(i, j int) bool { return candles[i] > candles[j] })

	var count int32
	a := []int32{}
	big := candles[0]

	for i := 0; i < len(candles); i++ {
		if candles[i] >= big {
			a = append(a, candles[i])
		}
	}

	count = int32(len(a))

	return count

}
