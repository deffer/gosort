package main

import (
	"strconv"
	"testing"
)

var initial = []int{5, 9, 12, 2, 6, -1, 200, 201, 10, -1000, 1000, 2}

type Order struct {
	number int
}

func (o Order) String() string {
	return strconv.Itoa(o.number)
}

func byNumber(one, two interface{}) int {
	o1, o2 := one.(Order), two.(Order) // SPARTAAAAAA!!!!!!

	if o1.number > o2.number {
		return 1
	} else if o1.number == o2.number {
		return 0
	} else {
		return -1
	}
}

func Test(t *testing.T) {
	toSort := make([]Order, 12)

	for i := range initial {
		toSort[i] = Order{initial[i]}
	}

	lessFunc := func(i, j int) bool {
		return toSort[i].number < toSort[j].number
	}
	SortStructs(toSort, lessFunc)

	for i := range toSort {
		if i > 0 {
			if byNumber(toSort[i-1], toSort[i]) == 1 {
				t.Errorf("Wrong sort, %v is bigger than %v", toSort[i-1], toSort[i])
			}
		}

	}
}
