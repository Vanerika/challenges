package main

import "fmt"

func main() {
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here

	//MANZANAS
	for i := 0; i < len(apples); i++ {
		apples[i] = apples[i] + a
	}

	//NARANJAS
	for i := 0; i < len(oranges); i++ {
		oranges[i] = oranges[i] + b
	}

	var count_apple int32
	var count_orange int32

	for _, apples := range apples {
		if apples >= s && apples <= t {
			count_apple++
		}
	}

	for _, oranges := range oranges {
		if oranges >= s && oranges <= t {
			count_orange++
		}
	}

	fmt.Println(count_apple)
	fmt.Println(count_orange)

}
