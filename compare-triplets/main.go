package main

import "fmt"

func main() {
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}
	fmt.Println(compareTriplets(a, b))
}

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	result := []int32{0, 0}
	for i, j := range a {
		switch {
		case j > b[i]:
			result[0]++
		case j < b[i]:
			result[1]++
		default:

		}
	}
	return result
}
