package sorting

import (
	"strconv"
	"testing"
)

var initial = []int{5, 9, 12, 2, 6, -1, 200, 201, 10, -1000, 1000, 2, 11}

type Order struct {
	number int
}

func (o Order) String() string {
	return strconv.Itoa(o.number)
}

func Test(t *testing.T) {
	orders := make([]Order, 13)

	for i := range initial {
		orders[i] = Order{initial[i]}
	}

	SortStructs(orders, func(i, j int) bool { // <-------------------
		return orders[i].number < orders[j].number
	})

	for i := range orders {
		if i > 0 {
			if orders[i-1].number > orders[i].number {
				t.Errorf("Wrong sort (structures), %v is bigger than %v", orders[i-1], orders[i])
			}
		}

	}
}
