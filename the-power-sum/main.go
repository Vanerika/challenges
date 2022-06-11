package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(powerSum(100, 2))
}

func powerSum(X int32, N int32) int32 {
	// Write your code here

	return newLimit(X, N, 1)

}

func newLimit(numero int32, pow int32, base int32) int32 {

	resto := numero - int32(math.Pow(float64(base), float64(pow)))

	if resto == 0 {
		return 1

	} else if resto < 0 {
		return 0
	} else {
		return (newLimit(resto, pow, base+1) + newLimit(numero, pow, base+1))
	}
}
