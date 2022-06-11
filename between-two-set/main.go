package main

import "fmt"

func main() {
	a := []int32{2, 3, 6}
	b := []int32{42, 84}

	fmt.Println(lcm(a))
	fmt.Println(gcd(b))

	fmt.Println(getTotalX(a, b))

}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here

	lcm := lcm(a)
	gcd := gcd(b)

	var count int32
	var i int32

	for i <= gcd {
		i = i + lcm
		if gcd%i == 0 {
			count++
			fmt.Println(i)
		}

	}

	return count

}

//GREATEST COMMON DIVISOR
func auxGCD(a, b int32) int32 {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func gcd(b []int32) int32 {
	var rdo int32 = b[0]
	for i := 1; i < len(b); i++ {
		rdo = auxGCD(rdo, b[i])
	}
	return rdo
}

func auxLCM(a, b int32) int32 {
	return a * (b / (auxGCD(a, b)))
}

func lcm(a []int32) int32 {
	var rdo int32 = a[0]
	for i := 1; i < len(a); i++ {
		rdo = auxLCM(rdo, a[i])
	}
	return rdo
}
