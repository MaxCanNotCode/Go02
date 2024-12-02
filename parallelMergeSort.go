package main

import (
	"sync"
)

func parallelDivide(left int, right int, slice []int32, aux []int32, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when this function returns

	if left >= right {
		return
	}

	middle := (left + right) / 2

	// WaitGroup for left and right subarrays
	var subWg sync.WaitGroup
	subWg.Add(2)

	// Left subarray
	select {
	case semaphore <- struct{}{}: // Acquire a slot
		// Left subarray in a new thread
		go func() {
			parallelDivide(left, middle, slice, aux, semaphore, &subWg)
			<-semaphore // Release the slot after finishing
		}()
	default:
		// Left subarray in the current thread (sequentially)
		parallelDivide(left, middle, slice, aux, semaphore, &subWg)
	}

	// Right subarray always processed in the current thread
	parallelDivide(middle+1, right, slice, aux, semaphore, &subWg)

	// Wait for both subarrays to finish sorting
	subWg.Wait()

	var mergeGroup sync.WaitGroup
	mergeGroup.Add(1)
	// Merge the sorted subarrays
	parallelConquer(left, middle, right, slice, aux, semaphore, &mergeGroup)
	mergeGroup.Wait()
}

func parallelConquer(left int, middle int, right int, slice []int32, aux []int32, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	var subWg sync.WaitGroup
	subWg.Add(2)

	// Process left and right subarrays in parallel
	select {
	case semaphore <- struct{}{}: // Acquire a slot
		go func() {
			mergeLeft(left, middle, middle+1, right, slice, aux, &subWg)
			<-semaphore
		}()
	default:
		mergeLeft(left, middle, middle+1, right, slice, aux, &subWg)
	}

	mergeRight(middle+1, right, left, middle, slice, aux, &subWg)

	subWg.Wait()

	// Copy merged result back to original slice
	if right+1-left >= 0 {
		copy(slice[left:right+1], aux[left:right+1])
	}
}

// mergeLeft handles elements from the left subarray
func mergeLeft(leftStart, leftEnd, rightStart, rightEnd int, slice []int32, aux []int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := leftStart; i <= leftEnd; i++ {
		val := slice[i]
		// Find position in right subarray
		pos := findPos(slice, rightStart, rightEnd, val)
		// Calculate and use final position
		finalPos := leftStart + (i - leftStart) + pos
		aux[finalPos] = val
	}
}

// mergeRight handles elements from the right subarray
func mergeRight(rightStart, rightEnd, leftStart, leftEnd int, slice []int32, aux []int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := rightStart; i <= rightEnd; i++ {
		val := slice[i]
		// Find position in left subarray
		pos := findPos(slice, leftStart, leftEnd, val)
		// Calculate and use final position
		finalPos := leftStart + pos + (i - rightStart)
		aux[finalPos] = val
	}
}

func findPos(slice []int32, start, end int, val int32) int {
	if start > end {
		return 0
	}

	left := start
	right := end + 1

	// Binary search to find the leftmost position where val should be inserted
	for left < right {
		mid := left + (right-left)/2
		if slice[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}

	//fmt.Printf("findPos: val=%d, pos=%d, start=%d, end=%d %v\n", val, left-start, start, end, slice[start:end+1])

	return left - start
}
