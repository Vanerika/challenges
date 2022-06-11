package main

import "fmt"

func main() {
	fmt.Println(catAndMouse(1, 2, 3))

}

func catAndMouse(x int32, y int32, z int32) string {

	var distancia_a int32 = x - z
	var distancia_b int32 = y - z

	if distancia_a < 0 {
		distancia_a = distancia_a * (-1)
	}
	if distancia_b < 0 {
		distancia_b = distancia_b * (-1)
	}

	if distancia_a > distancia_b {
		return "Cat B"
	} else if distancia_a < distancia_b {
		return "Cat A"
	} else {
		return "Mouse C"
	}

}
