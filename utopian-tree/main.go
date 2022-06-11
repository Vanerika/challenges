package main

import "fmt"

func main() {
	fmt.Println(utopianTree(4))
}

func utopianTree(n int32) int32 {
	// Write your code here

	var anterior int32 = 1
	var i int32
	for i = 1; i <= n; i++ {
		if i%2 == 0 {
			anterior = anterior + 1
		} else {
			anterior = anterior * 2
		}

	}

	return anterior

}
