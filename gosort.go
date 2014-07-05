package main

import (
	"fmt"
	"reflect"
	"sort"
)

type StructsSorter struct {
	originalSlice interface{}
	comparator    func(one, two interface{}) int
	l             int
	v             reflect.Value
	itemType      reflect.Type
}

func (s StructsSorter) getValueAt(i int) reflect.Value {
	return s.v.Index(i)
}

func (s StructsSorter) Len() int {
	return s.l
}

func (s StructsSorter) Swap(i, j int) {
	fi := s.getValueAt(i)
	fj := s.getValueAt(j)
	tmp := reflect.New(s.itemType).Elem()
	tmp.Set(fi)
	fi.Set(fj)
	fj.Set(tmp)

	fmt.Printf("Swapping %d with %d", i, j)
}
func (s StructsSorter) Less(i, j int) bool {
	return (s.comparator(s.getValueAt(i), s.getValueAt(j)) == -1)
}

func sorti(slice []interface{}, comparator func(one, two interface{}) int) {
	quicksort(0, len(slice)-1, slice, comparator)
}

func sorts(slice interface{}, comparator func(one, two interface{}) int) {
	sorter := StructsSorter{originalSlice: slice, comparator: comparator}
	sorter.v = reflect.ValueOf(slice)
	sorter.l = sorter.v.Len()
	sorter.itemType = sorter.getValueAt(0).Type()
	sort.Sort(sorter)
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
