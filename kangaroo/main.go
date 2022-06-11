package main

import "fmt"

func main() {
	fmt.Println(kangaroo(0, 2, 5, 3))
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here

	if (x2-x1)*(v2-v1) < 0 && (x2-x1)%(v2-v1) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
