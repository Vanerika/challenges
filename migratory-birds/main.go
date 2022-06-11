package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	migratoryBirds(arr)
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here

	birds := map[int32]int32{}
	var minor int32
	var that_bird int32 = math.MaxInt32

	for _, j := range arr {
		birds[j] = birds[j] + 1
	}

	for _, element := range birds {
		if element >= 2 && element > minor {
			minor = element
		}
	}

	for key, element := range birds {
		if element == minor && key < that_bird {
			that_bird = key
		}
	}

	fmt.Println(birds)
	fmt.Println(minor)
	fmt.Println(that_bird)
	return that_bird
}
