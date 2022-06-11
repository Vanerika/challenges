package main

import "fmt"

func main() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(ar))
}

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var result int64
	for _, j := range ar {
		result = result + j
	}
	return result
}
