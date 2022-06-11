package main

import "fmt"

func main() {
	var n int32
	n = 9
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	fmt.Println(sockMerchant(n, ar))

}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here

	var i int32

	m := make(map[int32]int32)

	for i = 0; i < n; i++ {
		m[ar[i]] = m[ar[i]] + 1
	}
	var count int32

	for _, element := range m {
		count = count + (element / 2)
	}

	return count
}
