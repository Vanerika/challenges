package main

import "fmt"

func main() {
	c := []int32{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(jumpingOnClouds(c))
}

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	jump := []int{}

	for i := 0; i < len(c); i++ {
		if len(c) > (i+2) && c[i+2] != 1 {
			jump = append(jump, i+2)
			i++
		} else if len(c) > (i+1) && c[i+1] != 1 {
			jump = append(jump, i+1)
		} else {
			break
		}

	}

	var count int32
	count = int32(len(jump))

	fmt.Println(jump)
	return count
}
