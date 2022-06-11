package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{769082435, 210437958, 673982045, 375809214, 380564127}

	miniMaxSum(arr)

}

func miniMaxSum(arr []int32) {
	// Write your code here

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var total int64

	for i := 0; i < 5; i++ {

		total = total + int64(arr[i])
	}
	fmt.Print((total - int64(arr[4])), " ", (total - int64(arr[0])))

}
