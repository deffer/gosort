package main

import (
	"fmt"
	"strconv"
)

var initial = []int{5, 9, 12, 2, 6}

type Order struct {
	number int
}

func (o Order) String() string {
	return strconv.Itoa(o.number)
}

func byNumber(one, two interface{}) int {
	o1, ok := one.(Order)
	if !ok {
		panic("Oops")
	}

	o2, ok := two.(Order)
	if !ok {
		panic("Oops")
	}

	if o1.number > o2.number {
		return 1
	} else if o1.number == o2.number {
		return 0
	} else {
		return -1
	}
}

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
	// Recursion
	if low < j {
		quicksort(low, j, slice, comparator)
	}
	if i < high {
		quicksort(i, high, slice, comparator)
	}
}

func main() {

	toSort := make([]interface{}, 5)

	for i := range initial {
		toSort[i] = Order{initial[i]}
	}

	sort(toSort, byNumber)

	for i := range toSort {
		fmt.Println(toSort[i])
	}
}
