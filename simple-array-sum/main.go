package main

import "fmt"

func main() {
	ar := []int32{1, 2, 3}
	fmt.Println(simpleArraySum(ar))

}

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32

	for _, j := range ar {
		sum = sum + j
	}
	return sum
}
