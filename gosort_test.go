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

func Test(t *testing.T) {
	toSort := make([]interface{}, 12)

	for i := range initial {
		toSort[i] = Order{initial[i]}
	}

	sort(toSort, byNumber)

	for i := range toSort {
		if i > 0 {
			if byNumber(toSort[i-1], toSort[i]) == 1 {
				t.Errorf("Wrong sort, %v is bigger than %v", toSort[i-1], toSort[i])
			}
		}

	}
}
