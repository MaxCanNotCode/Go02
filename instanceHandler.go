package main

import (
	"math"
	"math/rand"
)

func createSlice(exp int) (slice []int32) {
	size := int(math.Pow10(exp))
	slice = make([]int32, size)

	for i := 0; i < size; i++ {
		slice[i] = int32(i)
	}

	scramble(slice)
	return
}

/*
* Fisher-Yates-Shuffle
 */
func scramble(slice []int32) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func correctness(slice []int32) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] != int32(i) {
			return false
		}
	}

	return true
}
