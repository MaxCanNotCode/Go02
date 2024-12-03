package main

import (
	"sync"
)

func parallelQuickSort(slice []int32, start int, end int, p int, wg *sync.WaitGroup) {
	if start >= end {
		if wg != nil {
			wg.Done()
		}
		return
	}

	i := rearrange(start, end, slice)

	// We can only create a new goroutine if the number of active goroutines is below the limit
	if p > 1 {
		var localWg sync.WaitGroup
		localWg.Add(1)
		go func() {
			parallelQuickSort(slice, start, i+1, p/2, &localWg)
		}()
		parallelQuickSort(slice, i+2, end, p/2, nil)
		localWg.Wait()
	} else {
		parallelQuickSort(slice, start, i+1, 0, nil)
		parallelQuickSort(slice, i+2, end, 0, nil)
	}

	// Ensure the global WaitGroup is decremented if needed
	if wg != nil {
		wg.Done()
	}
}
