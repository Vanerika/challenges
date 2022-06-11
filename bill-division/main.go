package main

import (
	"fmt"
)

func main() {
	bill := []int32{3, 10, 2, 9}
	bonAppetit(bill, 1, 7)

}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here

	var anna_eat int32
	var total int32
	var i int32

	for i = 0; i < int32(len(bill)); i++ {
		if i != k {
			anna_eat = anna_eat + bill[i]
		}
		total = total + bill[i]
	}

	if anna_eat/2 == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(total/2 - anna_eat/2)
	}

}
