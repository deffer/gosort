package main

func sort(slice []interface{}, comparator func(one, two interface{}) int) {
	quicksort(0, len(slice)-1, slice, comparator)
}

func quicksort(low int, high int, slice []interface{}, comparator func(one, two interface{}) int) {
	i := low
	j := high

	pivot := slice[low+(high-low)/2]

	for i <= j {
		for lessThanPivot := comparator(slice[i], pivot); lessThanPivot == -1; lessThanPivot = comparator(slice[i], pivot) { // (numbers[i] < pivot) {
			i++
		}

		for moreThanPivot := comparator(slice[j], pivot); moreThanPivot == 1; moreThanPivot = comparator(slice[j], pivot) { //}   (numbers[j] > pivot) {
			j--
		}

		if i <= j {
			slice[i], slice[j] = slice[j], slice[i]
			i++
			j--
		}

	}

	if low < j {
		quicksort(low, j, slice, comparator)
	}
	if i < high {
		quicksort(i, high, slice, comparator)
	}
}
