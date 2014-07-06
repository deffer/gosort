package sorting

import (
	"reflect"
	"sort"
)

type StructsSorter struct {
	originalSlice interface{}
	lessFunc      func(i, j int) bool
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
}
func (s StructsSorter) Less(i, j int) bool {
	return s.lessFunc(i, j)
}

func SortStructs(slice interface{}, lessFunc func(i, j int) bool) {
	sorter := StructsSorter{originalSlice: slice, lessFunc: lessFunc}
	sorter.v = reflect.ValueOf(slice)
	sorter.l = sorter.v.Len()
	sorter.itemType = sorter.getValueAt(0).Type()
	sort.Sort(sorter)
}

/**
 * -------------------------------------------------------------
 * implementation of quick sort
 * -------------------------------------------------------------
 */
func Quicksort(slice []interface{}, comparator func(one, two interface{}) int) {
	quicksortr(0, len(slice)-1, slice, comparator)
}

func quicksortr(low int, high int, slice []interface{}, comparator func(one, two interface{}) int) {
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
		quicksortr(low, j, slice, comparator)
	}
	if i < high {
		quicksortr(i, high, slice, comparator)
	}
}
