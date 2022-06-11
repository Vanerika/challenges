package main

import "fmt"

func main() {
	s := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println(divisibleSumPairs(6, 5, s))

}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here

	var count int32

	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < (len(ar)); j++ {
			if i < j {
				pairs := []int32{ar[i], ar[j]}
				count = count + sum_array(pairs, k)
			}

		}
	}

	return count

}

func sum_array(arr []int32, k int32) int32 {

	var sum int32
	for _, j := range arr {
		sum = sum + j
	}
	if sum%k == 0 {
		return 1
	}
	return 0

}
