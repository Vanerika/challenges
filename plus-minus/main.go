package main

import "fmt"

func main() {

	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)

}

func plusMinus(arr []int32) {
	// Write your code here
	var plus float64
	var minus float64
	var zero float64

	for _, j := range arr {
		switch {
		case j > 0:
			plus++
		case j < 0:
			minus++
		default:
			zero++
		}
	}

	plus = plus / float64(len(arr))
	minus = minus / float64(len(arr))
	zero = zero / float64(len(arr))

	fmt.Println(plus)
	fmt.Println(minus)
	fmt.Println(zero)

}
