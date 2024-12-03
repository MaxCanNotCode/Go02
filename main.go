package main

import (
	"fmt"
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

		for j := 0; j < 1; j++ {
			var wg sync.WaitGroup
			wg.Add(1)
			p := i
			start := time.Now()
			go parallelQuickSort(slice, 0, len(slice), p, &wg, aux)
			wg.Wait()
			duration += time.Since(start).Milliseconds()
			correct := correctness(slice)
			if !correct {
				fmt.Println(slice)
				fmt.Println("Sorted incorrectly")
			}
			scramble(slice)
			copy(aux, slice)
		}
		fmt.Println("PARALLEL QUICKSORT With ", i, " Core(s): \n", duration, " ms total\n", duration/100, " ms average \n")

		/**
		for j := 0; j < 100; j++ {
			duration = callParMS(i, slice, aux, duration)
			scramble(slice)
			copy(aux, slice)
		}
		fmt.Println("PARALLEL MERGESORT With ", i, " Core(s): \n", duration, " ms total\n", duration/100, " ms average \n")
		duration = 0
		*/

	}
	/**
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
	*/

}

/**
func callParMS(i int, slice []int32, aux []int32, duration int64) int64 {
	semaphore := make(chan struct{}, i)
	var wg sync.WaitGroup
	wg.Add(1) // Top-level wait group for the whole divide-and-conquer process
	start := time.Now()
	go parallelDivide(0, len(slice)-1, slice, aux, semaphore, &wg)
	wg.Wait() // Wait for the entire sorting process to complete
	duration += time.Since(start).Milliseconds()
	correct := correctness(slice)
	if !correct {
		fmt.Println("Sorted incorrectly")
	}
	return duration
}
*/
