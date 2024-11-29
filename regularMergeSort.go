package main

func ms(slice []int32) {
	aux := make([]int32, len(slice))
	copy(aux, slice)
	divide(0, len(slice)-1, slice, aux)
}

func divide(left int, right int, slice []int32, aux []int32) {
	if left >= right {
		return
	}

	middle := (left + right) / 2

	divide(left, middle, slice, aux)
	divide(middle+1, right, slice, aux)

	conquer(left, middle, right, slice, aux)
}

func conquer(left, middle, right int, slice []int32, aux []int32) {
	i := left
	j := middle + 1
	k := left

	for i <= middle && j <= right {
		if slice[i] <= slice[j] {
			aux[k] = slice[i]
			i++
		} else {
			aux[k] = slice[j]
			j++
		}
		k++
	}

	for i <= middle {
		aux[k] = slice[i]
		i++
		k++
	}

	for j <= right {
		aux[k] = slice[j]
		j++
		k++
	}

	if right+1-left >= 0 {
		copy(slice[left:right+1], aux[left:right+1])
	}
}
