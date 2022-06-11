package main

import (
	"fmt"
)

func main() {

	//fmt.Println(factorial(5))

	sales := []int32{1, 2, 3, 4, 6}
	fmt.Println(balancedSales(sales))

	//fmt.Println(counterGame(6))

}

func factorial(n int32) int32 {

	var i int32
	var factorial int32 = 1

	for i = 1; i <= n; i++ {
		factorial = factorial * i
	}

	return factorial
}

func balancedSales(sales []int32) int32 {
	izquierda := []int32{}
	derecha := []int32{}

	var suma_izquierda int32
	var suma_derecha int32

	for i := 0; i < len(sales); i++ {
		suma_izquierda = suma_izquierda + sales[i]

		izquierda = append(izquierda, suma_izquierda)

		suma_derecha = suma_derecha + (sales[len(sales)-1-i])

		derecha = append(derecha, suma_derecha)

	}

	fmt.Println(izquierda)
	fmt.Println(derecha)
	for m := 0; m < len(sales); m++ {
		if izquierda[m] == derecha[len(sales)-1-m] {
			return int32(m)
		}
	}

	return 0
}

func counterGame(n int64) string {
	// Write your code here
	count := auxiliar(n, 0)

	fmt.Println(count)
	if count%2 == 0 {
		return "Louise"
	} else {
		return "Richard"
	}

}

func auxiliar(n int64, count int64) int64 {

	var base int64 = 1

	var i int64

	for i = 1; i <= n; i++ {
		if base < n && n != 1 {
			base = base * 2
		} else if base == n && n != 1 {
			n = n / 2
			return auxiliar(n, count+1)
		} else if n != 1 {
			n = n - (base / 2)
			return auxiliar(n, count+1)
		}

	}
	return count + 1
}
