package main

import "fmt"

func main() {
	ejemplo := []int32{1, 2, 2, 3, 1, 2}
	fmt.Println(pickingNumbers(ejemplo))
}

func pickingNumbers(a []int32) int32 {
	// Write your code here
	m := map[int32]int32{}
	var salida int32

	for _, j := range a {
		m[j] = m[j] + 1
	}

	for key := range m {
		if m[key]+m[key+1] > salida {
			salida = m[key] + m[key+1]
		}
	}

	return salida

}
