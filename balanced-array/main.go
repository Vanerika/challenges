package main

import "fmt"

func main() {
	array := []int32{1, 2, 1, 2, 1, 3}
	fmt.Println(balancedArray(array))

}

func balancedArray(sales []int32) int32 {
	//izquierda := []int32{}
	//derecha := []int32{}

	var suma_izquierda int32
	var suma_derecha int32

	var mas_peque int32

	for i := 0; i < len(sales)/2; i++ {
		suma_izquierda = suma_izquierda + sales[i]

		//izquierda = append(izquierda, sales[i])

		suma_derecha = suma_derecha + (sales[len(sales)-1-i])

		//derecha = append(derecha, sales[len(sales)-1-i])

	}

	if suma_derecha > suma_izquierda {
		mas_peque = suma_derecha - suma_izquierda
	} else if suma_derecha < suma_izquierda {
		mas_peque = suma_izquierda - suma_derecha
	} else {
		mas_peque = 0
	}

	return mas_peque
}
