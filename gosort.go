package main

import (
"fmt"
"strconv"
)

var initial = []int{5,9,12,2,6}

type Order struct{
	number int
}

func (o Order) String() string{
	return strconv.Itoa(o.number)
}

func byNumber(one, two interface{}) int{
	o1, ok := one.(Order)
	if !ok{
		panic("Oops")
	}
	
	o2, ok := two.(Order)
	if !ok{
		panic("Oops")
	}
	
	if o1.number > o2.number{
		return 1
	}else if o1.number == o2.number {
		return 0
	}else{
		return -1
	}
}

func sort(slice []interface{}, comparator func(one, two interface{}) int){
	for i := range slice{
		if (i>0){
			fmt.Printf("Comparing %v and %v: %d\n", slice[i], slice[i-1], 
			comparator(slice[i], slice[i-1]))
		}
	}
}

func main() {
	fmt.Printf("%0"+strconv.Itoa(9)+"d\n", 11)
	
	toSort := make([]interface{}, 5)
	
	for i:= range initial{
		toSort[i]=Order{initial[i]}
	}
	
	sort(toSort, byNumber)
}