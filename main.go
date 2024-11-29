package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {

	for i := 1; i <= runtime.NumCPU(); i++ {
		var duration int64 = 0
		slice := createSlice(6)

		aux := make([]int32, len(slice))
		copy(aux, slice)
		for j := 0; j < 100; j++ {
			start := time.Now()
			var wg sync.WaitGroup
			wg.Add(1)
			aux := make([]int32, len(slice))
			copy(aux, slice)
			go parallelDivide(0, len(slice)-1, slice, &wg, i, aux)
			wg.Wait()
			duration += time.Since(start).Milliseconds()
			correct := correctness(slice)
			if !correct {
				fmt.Println("!!!")
			}
			scramble(slice)
			copy(aux, slice)
		}
		fmt.Println("With ", i, " Core(s): \n", duration, " ms total\n", duration/100, " ms average \n")
	}

	for i := 3; i < 10; i++ {
		slice := createSlice(i)

		start := time.Now()
		qs(0, len(slice), slice)
		duration := time.Since(start)
		fmt.Print("Slice of size ", math.Pow10(i), " was sorted in ", duration, " using QuickSort ")
		correct := correctness(slice)
		fmt.Println("correct:", correct)

		scramble(slice)

		if i != 9 {
			start = time.Now()
			ms(slice)
			duration = time.Since(start)
			fmt.Print("Slice of size ", math.Pow10(i), " was sorted in ", duration, " using MergeSort ")
			correct = correctness(slice)
			fmt.Println("correct:", correct)
		}

		fmt.Println()
	}
}
