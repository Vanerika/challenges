package main

import "fmt"

func main() {
	staircase(4)
}

func staircase(n int32) {
	// Write your code here

	auxiliar(n, 1)

}

func auxiliar(n int32, x int32) {

	if x <= n {
		var i int32
		for i = 1; i <= (n - x); i++ {
			fmt.Print(" ")
		}
		for i = 1; i <= x; i++ {
			fmt.Print("#")
		}
		fmt.Println("")
		auxiliar(n, x+1)
	}
}
