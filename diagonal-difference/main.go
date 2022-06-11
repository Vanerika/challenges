package main

func main() {
	a := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	diagonalDifference(a)
}
func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var primary int32
	var secondary int32
	var diferencia int32

	//Primay
	for i := 0; i < len(arr); i++ {
		primary = primary + (arr[i][i])
	}

	//Secondary
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j >= 0; j-- {
			secondary = secondary + arr[i][j]
			i++
		}
	}

	diferencia = primary - secondary

	if diferencia < 0 {
		diferencia = diferencia * (-1)
	}

	return diferencia

}
