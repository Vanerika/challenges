package main

import (
	"fmt"
	"sort"
)

func main() {
	ranked := []int32{100, 100, 50, 40, 40, 20, 10}
	player := []int32{5, 25, 50, 120}
	fmt.Println(climbingLeaderboard(ranked, player))

}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here

	puestos := []int32{}
	//var puesto int32 = 1
	//var mayor int32 = math.MaxInt32

	nuevo_scores := []int32{}
	var anterior int32

	//m := map[int32]int32{}

	for _, j := range ranked {
		if j != anterior {
			anterior = j
			nuevo_scores = append(nuevo_scores, j)
		}
	}

	for _, m := range player {

		i := sort.Search(len(nuevo_scores), func(i int) bool { return nuevo_scores[i] <= m })
		puestos = append(puestos, int32(i)+1)
	}

	//nuevo_scores := []int32{}

	//sort.Slice(nuevo_scores, func(i, j int) bool { return nuevo_scores[i] > nuevo_scores[j] })

	// for _, j := range ranked {
	// 	if j != mayor && mayor > j {
	// 		mayor = j
	// 		m[j] = puesto
	// 		puesto = puesto + 1
	// 	}

	// }

	// fmt.Println(m)

	// for _, j := range player {
	// 	if m[j] != 0 {
	// 		puestos = append(puestos, m[j])
	// 	} else {
	// 		for key, element := range m {
	// 			puesto = 0
	// 			if key <= j {
	// 				puesto = element
	// 				break
	// 			} else if j < key {
	// 				puesto = element
	// 			}

	// 		}
	// 		puestos = append(puestos, puesto)

	// 	}

	// }

	return puestos

}
