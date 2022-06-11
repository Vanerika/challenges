package main

import "fmt"

func main() {

	var k int32 = 3
	a := []int32{-1, -3, -4, -2}
	fmt.Println(angryProfessor(k, a))
}

func angryProfessor(k int32, a []int32) string {
	// Write your code here

	var count_tempranos int32
	var salida string
	for _, j := range a {
		if j >= 0 {
			count_tempranos = count_tempranos + 1
		}
	}

	if count_tempranos >= k {
		salida = "NO"
	} else {
		salida = "NO"
	}

	return salida
}
