package main

import "sync"

func parallelQuickSort(slice []int32, start int, end int, p int, wg *sync.WaitGroup, aux []int32) {
	if start >= end-1 {
		if wg != nil {
			wg.Done()
		}
		return
	}

	if end-start >= p && p > 1 {
		smallerCounts := make([]int, p)
		biggerCounts := make([]int, p)
		localWg := sync.WaitGroup{}
		localWg.Add(p)
		segmentSize := (end - start) / p
		pivot := slice[end-1] // Use last element as pivot

		for i := 0; i < p; i++ {
			segStart := start + i*segmentSize
			segEnd := start + (i+1)*segmentSize
			if i == p-1 {
				segEnd = end - 1 // Exclude pivot for last segment
			}
			go parPartition(segStart, segEnd, slice, &localWg, smallerCounts, biggerCounts, i, pivot)
		}
		localWg.Wait()

		i := arrange(slice, smallerCounts, biggerCounts, start, end, aux, pivot)

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

func parPartition(start int, end int, slice []int32, wg *sync.WaitGroup, smallerCounts []int, biggerCounts []int, index int, pivot int32) {
	defer wg.Done()

	// Count elements smaller than pivot and partition the subarray
	smallerCount := 0
	biggerCount := 0
	i := start // Index for smaller elements

	for j := start; j < end; j++ {
		if slice[j] < pivot {
			// Swap current element with first larger element
			slice[i], slice[j] = slice[j], slice[i]
			i++
			smallerCount++
		} else {
			biggerCount++
		}
	}
	biggerCounts[index] = biggerCount
	smallerCounts[index] = smallerCount
}

func arrange(slice []int32, smallerCounts []int, biggerCounts []int, start int, end int, aux []int32, pivot int32) int {
	// Calculate pivot position relative to start
	pivotPos := start
	for i := 0; i < len(smallerCounts); i++ {
		pivotPos += smallerCounts[i]
	}

	// Place pivot in its final position

	// Copy elements to auxiliary array
	small := start
	big := pivotPos + 1
	cur := start

	for i := 0; i < len(smallerCounts); i++ {
		if smallerCounts[i] > 0 {
			copy(aux[small:small+smallerCounts[i]], slice[cur:cur+smallerCounts[i]])
			small += smallerCounts[i]
		}
		cur += smallerCounts[i]
		if biggerCounts[i] > 0 {
			copy(aux[big:big+biggerCounts[i]], slice[cur:cur+biggerCounts[i]])
			big += biggerCounts[i]
		}
		cur += biggerCounts[i]
	}

	slice[pivotPos] = pivot
	// Copy back to original slice
	copy(slice[start:pivotPos], aux[start:pivotPos]) // Copy smaller elements
	copy(slice[pivotPos+1:end], aux[pivotPos+1:end]) // Copy bigger elements

	return pivotPos
}

func seqQS(slice []int32, start int, end int) {
	if start >= end {
		return
	}
	i := partition(start, end, slice)
	seqQS(slice, start, i+1)
	seqQS(slice, i+2, end)
}
