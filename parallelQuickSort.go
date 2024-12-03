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
	// Choose middle pivot as the main pivot
	midPivotIdx := pivots[p/2]
	pivot := slice[midPivotIdx]

	// Arrays to store counts of elements smaller and larger than pivot
	smallerCounts := make([]int, p)
	largerCounts := make([]int, p)

	// Count smaller and larger elements in each partition
	segmentSize := (end - start) / p
	for i := 0; i < p; i++ {
		segStart := start + i*segmentSize
		segEnd := start + (i+1)*segmentSize
		if i == p-1 {
			segEnd = end
		}

		for j := segStart; j < segEnd; j++ {
			if slice[j] < pivot {
				smallerCounts[i]++
			} else if slice[j] > pivot {
				largerCounts[i]++
			}
		}
	}

	// Calculate prefix sums for position computation
	smallerPrefix := make([]int, p+1)
	largerPrefix := make([]int, p+1)

	for i := 0; i < p; i++ {
		smallerPrefix[i+1] = smallerPrefix[i] + smallerCounts[i]
		largerPrefix[i+1] = largerPrefix[i] + largerCounts[i]
	}

	equalPos := start + smallerPrefix[p]
	largerPos := end - largerPrefix[p]

	// Use auxiliary array for rearrangement
	copy(aux[start:end], slice[start:end])

	// Place elements in their final positions
	for i := 0; i < p; i++ {
		segStart := start + i*segmentSize
		segEnd := start + (i+1)*segmentSize
		if i == p-1 {
			segEnd = end
		}

		smallerStart := start + smallerPrefix[i]
		equalStart := start + smallerPrefix[p] + (i * ((largerPos - equalPos) / p))
		largerStart := end - largerPrefix[p] + largerPrefix[i]

		// Distribute elements to their respective positions
		for j := segStart; j < segEnd; j++ {
			if aux[j] < pivot {
				slice[smallerStart] = aux[j]
				smallerStart++
			} else if aux[j] > pivot {
				slice[largerStart] = aux[j]
				largerStart++
			} else {
				slice[equalStart] = aux[j]
				equalStart++
			}
		}
	}

	// Return the position of the chosen pivot (middle of equal elements)
	return start + smallerPrefix[p] + (largerPos-equalPos)/2
}

func parPartition(start int, end int, slice []int32, wg *sync.WaitGroup, pivots []int, index int) {
	defer wg.Done()
	pivot := slice[end-1]
	i := start - 1
	j := start

	for k := start; k < end-1; k++ {
		if slice[j] < pivot {
			i++
			tmpVal := slice[i]
			slice[i] = slice[j]
			slice[j] = tmpVal
		}
		j++
	}

	slice[end-1] = slice[i+1]
	slice[i+1] = pivot
	pivots[index] = i + 1
}
