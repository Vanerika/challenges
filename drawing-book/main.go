package main

import "fmt"

func main() {
	fmt.Println(pageCount(6, 2))
}

func pageCount(n int32, p int32) int32 {
	// Write your code here
	var contando_pagina_cero int32 = n + 1
	var desde_el_frente int32
	var desde_atras int32
	if contando_pagina_cero%2 == 0 {
		desde_el_frente = p / 2
		desde_atras = (n - p) / 2

	} else {
		desde_el_frente = p / 2
		desde_atras = (n - p + 1) / 2
	}

	if desde_atras >= desde_el_frente {
		return desde_el_frente
	} else {
		return desde_atras
	}
}
