package main

import "fmt"

func main() {
	keyboard := []int32{40, 50, 60}
	drives := []int32{5, 8, 12}

	fmt.Println(getMoneySpent(keyboard, drives, 60))

}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {

	compra := []int32{}
	var gasto int32

	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			compra = append(compra, (keyboards[i] + drives[j]))
		}
	}

	for i := 0; i < len(compra); i++ {
		if compra[i] >= gasto && compra[i] <= b {
			gasto = compra[i]
		}
	}

	if gasto == 0 {
		return -1
	} else {
		return gasto
	}

}
