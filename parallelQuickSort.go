package main

import "sync"

func parallelQuickSort(slice []int32, start int, end int, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if start == end {
		return
	}

	i := rearrange(start, end, slice)

	var subWg sync.WaitGroup
	subWg.Add(2)

	select {
	case semaphore <- struct{}{}:
		go func() {
			parallelQuickSort(slice, start, i+1, semaphore, &subWg)
			<-semaphore
		}()
	default:
		parallelQuickSort(slice, start, i+1, semaphore, &subWg)
	}

	parallelQuickSort(slice, i+2, end, semaphore, &subWg)

	subWg.Wait()
}
