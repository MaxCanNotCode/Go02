package main

import "sync"

func parallelDivide(left int, right int, slice []int32, wg *sync.WaitGroup, p int, aux []int32) {
	if left >= right {
		if wg != nil {
			wg.Done()
		}
		return
	}

	middle := (left + right) / 2

	if p > 1 {
		var localWg sync.WaitGroup
		localWg.Add(1)
		go func() {
			parallelDivide(left, middle, slice, &localWg, p/2, aux)
		}()
		parallelDivide(middle+1, right, slice, nil, p/2, aux)

		localWg.Wait()
	} else {
		parallelDivide(left, middle, slice, nil, 0, aux)
		parallelDivide(middle+1, right, slice, nil, 0, aux)
	}

	conquer(left, middle, right, slice, aux)

	if wg != nil {
		wg.Done() // Ensure completion signal for the caller
	}
}
