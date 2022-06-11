package main

import (
	"fmt"
)

func main() {

	ejemplo := [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}}
	fmt.Println(formingMagicSquare(ejemplo))
}

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	combinaciones := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}}}

	totales := []int32{0, 0, 0, 0, 0, 0, 0, 0}
	var salida int32 = 200
	var a int32
	var b int32

	for i := 0; i < len(combinaciones); i++ {
		for j := 0; j < len(combinaciones[i]); j++ {
			for k := 0; k < len(combinaciones[i][j]); k++ {
				b = combinaciones[i][j][k]
				a = s[j][k]

				suma := a - b
				if suma < 0 {
					suma = suma * (-1)
				}

				totales[i] = totales[i] + suma
			}
		}
	}

	for _, i := range totales {
		if i < salida {
			salida = i
		}
	}

	return salida
}
