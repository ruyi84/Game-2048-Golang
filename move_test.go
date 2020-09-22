package main

import (
	"testing"
)

func Test_LeftMove(t *testing.T) {
	ints := []int{0, 2, 4}
	num := 1
	for num != 0 {
		num = 0
		for k := 0; k < len(ints)-1; k++ {
			if ints[k] == 0 && ints[k+1] != 0 {
				ints[k], ints[k+1] = ints[k+1], ints[k]
				num++
				continue
			}
			if ints[k] == ints[k+1] && ints[k] != 0 {
				ints[k] += ints[k+1]
				ints[k+1] = 0
				num++
				continue
			}

		}
	}
}

func Test_reverse(t *testing.T) {
	ints := [][]int{[]int{0, 2, 4},[]int{0, 2, 4},[]int{0, 2, 4}}

	var newPlate [3][3]int

	num := len(ints) - 1
	for k:= range newPlate {

		for i := range newPlate[k] {
			newPlate[k][i] = ints[i][num]
		}
		num--
	}


}
