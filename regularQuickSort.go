package main

func qs(start int, end int, slice []int32) {
	if start == end {
		return
	}

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

	qs(start, i+1, slice)
	qs(i+2, end, slice)
}
