package main

import "fmt"

func main() {
	ejemplo := []int32{2, 5, 4, 5, 2}
	fmt.Println(hurdleRace(7, ejemplo))

}

func buscarMayor(a []int32) int32 {

	var numeroMayor int32 = 0
	for _, numero := range a {
		if numero > numeroMayor {
			numeroMayor = numero
		}
	}

	return numeroMayor

}

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here

	posiones := buscarMayor(height) - k

	if posiones <= 0 {
		return 0
	} else {
		return posiones
	}

}
