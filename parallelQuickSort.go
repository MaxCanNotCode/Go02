package main

import (
	"sync"
)

func parallelQuickSort(slice []int32, start int, end int, p int, wg *sync.WaitGroup, aux []int32) {
	if start >= end-1 {
		if wg != nil {
			wg.Done()
		}
		return
	}

	if end-start >= p && p > 1 {
		pivots := make([]int, p)
		localWg := sync.WaitGroup{}
		localWg.Add(p)
		segmentSize := (end - start) / p

		for i := 0; i < p; i++ {
			segStart := start + i*segmentSize
			segEnd := start + (i+1)*segmentSize
			if i == p-1 {
				segEnd = end
			}
			go parPartition(segStart, segEnd, slice, &localWg, pivots, i)
		}
		localWg.Wait()

		i := arrange(slice, pivots, start, end, p, aux)

		// Sequential recursive calls
		parallelQuickSort(slice, start, i, p, nil, aux)
		parallelQuickSort(slice, i+1, end, p, nil, aux)
	} else {
		seqQS(slice, start, end)
	}

	if wg != nil {
		wg.Done()
	}
}

func seqQS(slice []int32, start int, end int) {
	if start >= end {
		return
	}
	i := partition(start, end, slice)
	seqQS(slice, start, i+1)
	seqQS(slice, i+2, end)
}

func arrange(slice []int32, pivots []int, start int, end int, p int, aux []int32) int {

}

func parPartition(start int, end int, slice []int32, wg *sync.WaitGroup, pivots []int, index int) {
	defer wg.Done()

}
