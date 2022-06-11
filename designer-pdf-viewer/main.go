package main

import "fmt"

func main() {
	a := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5}

	fmt.Println(designerPdfViewer(a, "torn"))
}

func turnToNumber(a string) []int32 {

	m := map[string]int32{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}

	salida := []int32{}

	for _, j := range a {
		salida = append(salida, m[string(rune(j))])
	}

	return salida
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

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here

	letras_a_numero := turnToNumber(word)

	alturas := []int32{}

	for _, j := range letras_a_numero {
		alturas = append(alturas, h[j])
	}

	el_mayor := buscarMayor(alturas)

	return el_mayor * int32(len(letras_a_numero))

}
