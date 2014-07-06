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
