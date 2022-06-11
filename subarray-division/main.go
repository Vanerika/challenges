package main

import "fmt"

func main() {
	s := []int32{4}
	fmt.Println(birthday(s, 4, 1))

}

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here

	sub_array := []int32{}
	var count int32

	for i := 0; i < len(s); i++ {
		if i+int(m) <= len(s) {
			sub_array = s[i : i+int(m)]
			fmt.Println(sub_array)
			count = count + sum_array(sub_array, d)
		}

	}

	return count
}

func sum_array(s []int32, d int32) int32 {

	var sum int32
	for _, j := range s {
		sum = sum + j
	}
	if sum == d {
		return 1
	}
	return 0

}
